/**
 * @author jiangshangfang
 * @date 2022/1/14 3:09 PM
 **/
package consul

import (
	"context"
	"fmt"
	"gin/pkg/registry"
	"github.com/hashicorp/consul/api"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	_ registry.Registry  = &Registry{}
	_ registry.Discovery = &Registry{}
)

// Registry is consul registry
type Registry struct {
	client            *api.Client
	ctx               context.Context
	cancel            context.CancelFunc
	enableHealthCheck bool
	registry          map[string]*serviceSet
	lock              sync.RWMutex
}

func New(client *api.Client, opts ...Option) *Registry {
	r := &Registry{
		client:            client,
		enableHealthCheck: true,
		registry:          make(map[string]*serviceSet),
	}
	for _, opt := range opts {
		opt(r)
	}
	r.ctx, r.cancel = context.WithCancel(context.Background())
	return r
}

func (r *Registry) Register(ctx context.Context, svc *registry.ServiceInstance) error {
	addresses := make(map[string]api.ServiceAddress)
	var addr string
	var port uint64
	for _, endpoint := range svc.Endpoints {
		raw, err := url.Parse(endpoint)
		if err != nil {
			return err
		}
		addr = raw.Hostname()
		port, _ = strconv.ParseUint(raw.Port(), 10, 16)
		addresses[raw.Scheme] = api.ServiceAddress{Address: endpoint, Port: int(port)}
	}
	asr := &api.AgentServiceRegistration{
		ID:              svc.ID,
		Name:            svc.Name,
		Meta:            svc.Metadata,
		Tags:            []string{fmt.Sprintf("version=%s", svc.Version)},
		TaggedAddresses: addresses,
		Address:         addr,
		Port:            int(port),
		Checks: []*api.AgentServiceCheck{
			{
				CheckID:                        "service:" + svc.ID,
				TTL:                            "30s",
				Status:                         "passing",
				DeregisterCriticalServiceAfter: "90s",
			},
		},
	}
	if r.enableHealthCheck {
		asr.Checks = append(asr.Checks, &api.AgentServiceCheck{
			TCP:                            fmt.Sprintf("%s:%d", addr, port),
			Interval:                       "20s",
			Status:                         "passing",
			DeregisterCriticalServiceAfter: "90s",
		})
	}
	err := r.client.Agent().ServiceRegister(asr)
	if err != nil {
		return err
	}
	go func() {
		ticker := time.NewTicker(20 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				_ = r.client.Agent().UpdateTTL("service:"+svc.ID, "pass", "pass")
			case <-r.ctx.Done():
				return
			}
		}
	}()
	return nil
}

// DeRegister the registration.
func (r *Registry) DeRegister(ctx context.Context, svc *registry.ServiceInstance) error {
	r.cancel()
	return r.client.Agent().ServiceDeregister(svc.ID)
}

func (r *Registry) GetService(ctx context.Context, name string) (services []*registry.ServiceInstance, err error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	set := r.registry[name]
	if set == nil {
		return nil, fmt.Errorf("service %s not resolved in registry", name)
	}
	ss, _ := set.services.Load().([]*registry.ServiceInstance)
	if ss == nil {
		return nil, fmt.Errorf("service %s not found in registry", name)
	}
	services = append(services, ss...)
	return
}

func (r *Registry) Watch(ctx context.Context, name string) (registry.Watcher, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	set, ok := r.registry[name]
	if !ok {
		set = &serviceSet{
			serviceName: name,
			watcher:     make(map[*watcher]struct{}),
			services:    &atomic.Value{},
		}
		r.registry[name] = set
	}
	// 初始化watcher
	w := &watcher{
		event: make(chan struct{}, 1),
	}
	w.ctx, w.cancel = context.WithCancel(context.Background())
	w.set = set
	set.lock.Lock()
	set.watcher[w] = struct{}{}
	set.lock.Unlock()
	ss, _ := set.services.Load().([]*registry.ServiceInstance)
	if len(ss) > 0 {
		// If the service has a value, it needs to be pushed to the watcher,
		// otherwise the initial data may be blocked forever during the watch.
		w.event <- struct{}{}
	}

	if !ok {
		go func(set *serviceSet) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			services, idx, err := r.service(ctx, set.serviceName, 0, true)
			cancel()
			if err == nil && len(services) > 0 {
				set.broadcast(services)
			}
			ticker := time.NewTicker(time.Second)
			defer ticker.Stop()
			for {
				<-ticker.C
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
				tmpService, tmpIdx, err := r.service(ctx, set.serviceName, idx, true)
				cancel()
				if err != nil {
					time.Sleep(time.Second)
					continue
				}
				if len(tmpService) != 0 && tmpIdx != idx {
					services = tmpService
					set.broadcast(services)
				}
				idx = tmpIdx
			}
		}(set)
	}
	return w, nil
}

func (r *Registry) service(ctx context.Context, service string, index uint64, passingOnly bool) ([]*registry.ServiceInstance, uint64, error) {
	opts := &api.QueryOptions{
		WaitIndex: index,
		WaitTime:  time.Second * 55,
	}
	opts = opts.WithContext(ctx)
	entries, meta, err := r.client.Health().Service(service, "", passingOnly, opts)
	if err != nil {
		return nil, 0, err
	}

	services := make([]*registry.ServiceInstance, 0)

	for _, entry := range entries {
		var version string
		for _, tag := range entry.Service.Tags {
			strs := strings.SplitN(tag, "=", 2)
			if len(strs) == 2 && strs[0] == "version" {
				version = strs[1]
			}
		}
		var endpoints []string
		for scheme, addr := range entry.Service.TaggedAddresses {
			if scheme == "lan_ipv4" || scheme == "wan_ipv4" || scheme == "lan_ipv6" || scheme == "wan_ipv6" {
				continue
			}
			endpoints = append(endpoints, addr.Address)
		}
		services = append(services, &registry.ServiceInstance{
			ID:        entry.Service.ID,
			Name:      entry.Service.Service,
			Metadata:  entry.Service.Meta,
			Version:   version,
			Endpoints: endpoints,
		})
	}
	return services, meta.LastIndex, nil
}

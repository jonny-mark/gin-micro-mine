/**
 * @author jiangshangfang
 * @date 2022/1/10 2:27 PM
 **/
package etcd

import (
	"context"
	"fmt"
	"gin-micro-mine/pkg/registry"
	"go.etcd.io/etcd/client/v3"
	"math/rand"
	"time"
)

var (
	_ registry.Registry  = &Registry{}
	_ registry.Discovery = &Registry{}
)

// Registry is etcd registry.
type Registry struct {
	opts   *options
	client *clientv3.Client
}

// New create a etcd registry
func New(client *clientv3.Client, opts ...Option) (r *Registry) {
	o := &options{
		ctx:       context.Background(),
		namespace: "/microServices",
		ttl:       time.Second * 15,
		maxRetry:  5,
	}
	for _, opt := range opts {
		opt(o)
	}
	return &Registry{
		opts:   o,
		client: client,
	}
}

func (r *Registry) Register(ctx context.Context, service *registry.ServiceInstance) error {
	key := fmt.Sprintf("%s/%s/%s", r.opts.namespace, service.Name, service.ID)
	value, err := marshal(service)
	if err != nil {
		return err
	}
	leaseID, err := r.registerWithKV(ctx, key, value)
	if err != nil {
		return err
	}

	go r.heartBeat(r.opts.ctx, leaseID, key, value)
	return nil
}

func (r *Registry) DeRegister(ctx context.Context, service *registry.ServiceInstance) error {
	key := fmt.Sprintf("%s/%s/%s", r.opts.namespace, service.Name, service.ID)
	_, err := r.client.Delete(ctx, key)
	return err
}

func (r *Registry) GetService(ctx context.Context, name string) ([]*registry.ServiceInstance, error) {
	key := fmt.Sprintf("%s/%s", r.opts.namespace, name)
	services, err := r.client.KV.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	items := make([]*registry.ServiceInstance, 0, len(services.Kvs))
	for _, kv := range services.Kvs {
		si, err := unmarshal(kv.Value)
		if err != nil {
			return nil, err
		}
		if si.Name != name {
			continue
		}
		items = append(items, si)
	}
	return items, nil
}

func (r *Registry) Watch(ctx context.Context, name string) (registry.Watcher, error) {
	key := fmt.Sprintf("%s/%s", r.opts.namespace, name)
	return newWatcher(ctx, key, name, r.client)
}

func (r *Registry) registerWithKV(ctx context.Context, key string, value string) (clientv3.LeaseID, error) {
	grant, err := r.client.Lease.Grant(ctx, int64(r.opts.ttl.Seconds()))
	if err != nil {
		return 0, err
	}
	_, err = r.client.Put(ctx, key, value, clientv3.WithLease(grant.ID))
	if err != nil {
		return 0, err
	}
	return grant.ID, nil
}

func (r *Registry) heartBeat(ctx context.Context, leaseID clientv3.LeaseID, key string, value string) {
	curLeaseID := leaseID

	kac, err := r.client.KeepAlive(ctx, leaseID)
	if err != nil {
		curLeaseID = 0
	}
	rand.Seed(time.Now().Unix())

	for {
		if curLeaseID == 0 {
			// try to registerWithKV
			for retryCnt := 0; retryCnt < r.opts.maxRetry; retryCnt++ {
				if ctx.Err() != nil {
					return
				}
				idChan := make(chan clientv3.LeaseID, 1)
				errChan := make(chan error, 1)
				cancelCtx, cancel := context.WithCancel(ctx)
				go func() {
					defer cancel()
					id, registerErr := r.registerWithKV(cancelCtx, key, value)
					if registerErr != nil {
						errChan <- registerErr
					} else {
						idChan <- id
					}
				}()

				//拿到新注册的LeaseID
				select {
				//超时推出
				case <-time.After(3 * time.Second):
					cancel()
					continue
				case <-errChan:
					continue
				case curLeaseID = <-idChan:
				}
				kac, err = r.client.KeepAlive(ctx, curLeaseID)
				if err == nil {
					break
				}
				time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			}
		}
		select {
		case _, ok := <-kac:
			if !ok {
				if ctx.Err() != nil {
					return
				}
				curLeaseID = 0
				continue
			}
		case <-r.opts.ctx.Done():
			return
		}
	}

}

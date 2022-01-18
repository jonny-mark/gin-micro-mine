/**
 * @author jiangshangfang
 * @date 2021/12/12 8:11 PM
 **/
package app

import (
	"context"
	"gin/pkg/log"
	"os"
	"syscall"
	"time"
	"github.com/google/uuid"
	"sync"
	"gin/pkg/registry"
	"gin/pkg/transport"
	"golang.org/x/sync/errgroup"
)

// App global app
type App struct {
	opts     options
	ctx      context.Context
	cancel   func()
	mu       sync.Mutex
	instance *registry.ServiceInstance
}

// New create a app globally
func New(opts ...Option) *App {
	o := options{
		ctx:    context.Background(),
		logger: log.GetLogger(),
		// don not catch SIGKILL signal, need to waiting for kill self by other.
		sigs:            []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
		registryTimeout: 10 * time.Second,
	}
	if id, err := uuid.NewUUID(); err == nil {
		o.id = id.String()
	}
	for _, opt := range opts {
		opt(&o)
	}

	ctx, cancel := context.WithCancel(o.ctx)
	return &App{
		opts:   o,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (a *App) Run() error {
	//获取instance，做服务注册
	instance, err := a.buildInstance()
	if err != nil {
		return err
	}

	//errgroup 确保所有服务正常启动
	group, errCtx := errgroup.WithContext(a.ctx)

	//同步启动
	wg := sync.WaitGroup{}
	for _, srv := range a.opts.servers {
		group.Go(func() error {
			//错误信号
			<-errCtx.Done()
			return srv.Stop(errCtx)
		})
		wg.Add(1)
		group.Go(func() error {
			wg.Done()
			return srv.Start(errCtx)
		})
	}

	wg.Wait()

}

func (a *App) buildInstance() (*registry.ServiceInstance, error) {
	endpoints := make([]string, 0)
	for _, endpoint := range a.opts.endpoints {
		endpoints = append(endpoints, endpoint.String())
	}
	//没有节点，注册instance
	if len(endpoints) == 0 {
		for _, srv := range a.opts.servers {
			if r, ok := srv.(transport.Endpoint); ok {
				e, err := r.Endpoint()
				if err != nil {
					return nil, err
				}
				endpoints = append(endpoints, e.String())
			}
		}
	}
	return &registry.ServiceInstance{
		ID:        a.opts.id,
		Name:      a.opts.name,
		Version:   a.opts.version,
		Metadata:  a.opts.metadata,
		Endpoints: endpoints,
	}, nil
}

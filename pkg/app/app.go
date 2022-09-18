/**
 * @author jiangshangfang
 * @date 2021/12/12 8:11 PM
 **/
package app

import (
	"context"
	"errors"
	logger "gin-micro-mine/pkg/log"
	"gin-micro-mine/pkg/registry"
	"gin-micro-mine/pkg/transport"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
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
		logger: logger.GetLogger(),
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

	//服务注册
	if a.opts.registry != nil {
		//获取instance
		instance, err := a.buildInstance()
		if err != nil {
			return err
		}
		c, cancel := context.WithTimeout(a.opts.ctx, a.opts.registryTimeout)
		defer cancel()
		if err = a.opts.registry.Register(c, instance); err != nil {
			return err
		}
		a.mu.Lock()
		a.instance = instance
		a.mu.Unlock()
	}

	//监听signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, a.opts.sigs...)
	group.Go(func() error {
		for {
			select {
			//监听errgroup的chan
			case <-errCtx.Done():
				return errCtx.Err()
			//监听signal信息
			case s := <-quit:
				a.opts.logger.Infof("receive a quit signal: %s", s.String())
				err := a.stop()
				if err != nil {
					a.opts.logger.Infof("failed to stop app, err: %s", err.Error())
					return err
				}
			}
		}
	})

	if err := group.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
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

//优雅关闭应用
func (a *App) stop() error {
	//如果开启注册，需要注销节点
	a.mu.Lock()
	instance := a.instance
	a.mu.Unlock()

	if a.opts.registry != nil && instance != nil {
		ctx, cancel := context.WithTimeout(a.ctx, a.opts.registryTimeout)
		defer cancel()
		if err := a.opts.registry.DeRegister(ctx, instance); err != nil {
			return err
		}
	}

	if a.cancel != nil {
		a.cancel()
	}
	return nil
}

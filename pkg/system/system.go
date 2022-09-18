/**
 * @author jiangshangfang
 * @date 2021/10/25 10:19 AM
 **/
package system

import (
	"context"
	"errors"
	"gin-micro-mine/pkg/config"
	"gin-micro-mine/pkg/log"
	"gin-micro-mine/pkg/transport"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

type System struct {
	name    string
	version string

	signals []os.Signal
	ctx     context.Context
	logger  log.Logger
	servers []transport.Server
	cancel  func()
	c       *config.Config
}

func New(cfg *config.Config, opts ...Options) *System {
	system := &System{
		name:    "gin",
		version: "1",
		signals: []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
		ctx:     context.Background(),
		logger:  log.GetLogger(),
	}

	for _, o := range opts {
		o(system)
	}
	_, cancel := context.WithCancel(system.ctx)
	system.c = cfg
	system.cancel = cancel
	return system
}

func (s *System) Run() error {
	s.logger.Infof("name: %s, version: %s", s.name, s.version)
	//请求中只要有一个有错误就会返回error
	eg, ctx := errgroup.WithContext(s.ctx)

	//start server
	for _, srv := range s.servers {
		eg.Go(func() error {
			<-ctx.Done()
			return srv.Stop(ctx)
		})
		eg.Go(func() error {
			return srv.Start(ctx)
		})
	}

	// watch signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, s.signals...)
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case q := <-quit:
				s.logger.Infof("receive a quit signal: %s", q.String())
				return s.Stop()
			}
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

// Stop stops the application gracefully.
func (s *System) Stop() error {
	if s.cancel != nil {
		s.cancel()
	}
	return nil
}

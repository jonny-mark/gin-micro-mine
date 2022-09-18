/**
 * @author jiangshangfang
 * @date 2021/10/24 9:43 PM
 **/
package system

import (
	"context"
	"github.com/jonny-mark/gin-micro-mine/pkg/log"
	"github.com/jonny-mark/gin-micro-mine/pkg/transport"
	"os"
)

type Options func(s *System)

// WithName .
func WithName(name string) Options {
	return func(s *System) {
		s.name = name
	}
}

// WithVersion with a version
func WithVersion(version string) Options {
	return func(s *System) {
		s.version = version
	}
}

// WithContext with a context
func WithContext(ctx context.Context) Options {
	return func(s *System) {
		s.ctx = ctx
	}
}

// WithSignal with some system signal
func WithSignal(sigs ...os.Signal) Options {
	return func(s *System) {
		s.signals = sigs
	}
}

// WithLogger .
func WithLogger(logger log.Logger) Options {
	return func(s *System) {
		s.logger = logger
	}
}

// WithServer with a server , http or grpc
func WithServer(srv ...transport.Server) Options {
	return func(s *System) {
		s.servers = srv
	}
}

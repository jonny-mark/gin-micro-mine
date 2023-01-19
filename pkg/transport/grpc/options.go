package grpc

import (
	"google.golang.org/grpc"
	"time"
)

type ServerOption func(s *Server)

func WithOptions(opts ...grpc.ServerOption) ServerOption {
	return func(s *Server) {
		s.grpcOpts = opts
	}
}

func WithInters(inters ...grpc.UnaryServerInterceptor) ServerOption {
	return func(s *Server) {
		s.inters = inters
	}
}

func WithNetwork(network string) ServerOption {
	return func(s *Server) {
		s.network = network
	}
}

func WithAddr(addr string) ServerOption {
	return func(s *Server) {
		s.addr = addr
	}
}

func WithReadTimeout(readTimeout time.Duration) ServerOption {
	return func(s *Server) {
		s.readTimeout = readTimeout
	}
}

func WithWriteTimeout(writeTimeout time.Duration) ServerOption {
	return func(s *Server) {
		s.writeTimeout = writeTimeout
	}
}

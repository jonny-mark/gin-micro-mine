/**
 * @author jiangshangfang
 * @date 2021/10/23 10:23 PM
 **/
package grpc

import (
	"time"
	"context"
	"net"
	"gin/pkg/log"
	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server
	ctx     context.Context
	lis     net.Listener
	timeout time.Duration
	network string
	address string
	inters  []grpc.UnaryServerInterceptor
	log     log.Logger
}

// NewServer creates a gRPC server by options.
func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		network: "tcp",
		address: ":0",
		timeout: 1 * time.Second,
		log:     log.GetLogger(),
	}
	for _, o := range opts {
		o(srv)
	}
	var ints = []grpc.UnaryServerInterceptor{}
	if len(srv.inters) > 0 {
		ints = append(ints, srv.inters...)
	}
	var grpcOpts = []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(ints...),
	}
	srv.Server = grpc.NewServer(grpcOpts...)

	return srv
}

// Start start the gRPC server
func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen(s.network, s.address)
	if err != nil {
		return err
	}
	s.lis = lis
	s.ctx = ctx
	s.log.Infof("[gRPC] server is listening on: %s", lis.Addr().String())
	if err := s.Serve(lis) ;err != nil{
		return err
	}
	return nil
}

// Stop stop the gRPC server.
func (s *Server) Stop(ctx context.Context) error {
	s.GracefulStop()
	s.log.Info("[gRPC] server is stopping")
	return nil
}
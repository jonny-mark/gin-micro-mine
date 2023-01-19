package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/url"
	"time"

	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/jonny-mark/gin-micro-mine/pkg/utils"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc/health"
	healthPb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	*grpc.Server
	grpcOpts     []grpc.ServerOption
	network      string
	addr         string
	readTimeout  time.Duration
	writeTimeout time.Duration
	inters       []grpc.UnaryServerInterceptor
	health       *health.Server
	listener     net.Listener
	endpoint     *url.URL
	ctx          context.Context

	EnableTracing bool
	TracerOptions []otelgrpc.Option
}

// NewServer creates a gRPC server by options.
func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		network:      "tcp",
		addr:         ":0",
		readTimeout:  5 * time.Second,
		writeTimeout: 5 * time.Second,
		health:       health.NewServer(),
	}

	for _, o := range opts {
		o(srv)
	}

	// 一元拦截器   构造拦截器链
	chainUnaryInterceptors := []grpc.UnaryServerInterceptor{
		unaryServerInterceptor(),
		grpcPrometheus.UnaryServerInterceptor,
		grpcRecovery.UnaryServerInterceptor(),
	}
	if len(srv.inters) > 0 {
		chainUnaryInterceptors = append(chainUnaryInterceptors, srv.inters...)
	}

	// 流试拦截器  构造拦截器链
	chainStreamInterceptors := []grpc.StreamServerInterceptor{
		grpcPrometheus.StreamServerInterceptor,
		grpcRecovery.StreamServerInterceptor(),
	}

	if srv.EnableTracing {
		chainUnaryInterceptors = append(chainUnaryInterceptors, otelgrpc.UnaryServerInterceptor(srv.TracerOptions...))
		chainStreamInterceptors = append(chainStreamInterceptors, otelgrpc.StreamServerInterceptor(srv.TracerOptions...))
	}

	grpcOpts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(chainUnaryInterceptors...),
		grpc.ChainStreamInterceptor(chainStreamInterceptors...),
	}

	if len(srv.grpcOpts) > 0 {
		grpcOpts = append(grpcOpts, srv.grpcOpts...)
	}
	grpcServer := grpc.NewServer(grpcOpts...)

	//grpc健康检查 参考 https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	srv.health.SetServingStatus("", healthPb.HealthCheckResponse_SERVING)
	healthPb.RegisterHealthServer(grpcServer, srv.health)
	reflection.Register(grpcServer)

	//注册grpcServer的零值指标
	grpcPrometheus.Register(grpcServer)

	srv.Server = grpcServer
	return srv
}

// Start start the gRPC server
func (s *Server) Start(ctx context.Context) error {
	listener, err := net.Listen(s.network, s.addr)
	if err != nil {
		log.Printf("[gRPC] server listen fail,network is %s,add is %s,rerr is %s", s.network, s.addr, err.Error())
		return err
	}
	s.listener = listener

	if _, err := s.Endpoint(); err != nil {
		log.Printf("[gRPC] server with endpoint fail,err is %s", err.Error())
		return err
	}
	s.ctx = ctx
	log.Printf("[gRPC] server is listening on: %s", s.listener.Addr().String())
	return s.Serve(s.listener)
}

// Stop stop the gRPC server.
func (s *Server) Stop(ctx context.Context) error {
	s.GracefulStop()
	log.Printf("[gRPC] server is stopping")
	return nil
}

// grpc://127.0.0.1:9090
func (s *Server) Endpoint() (*url.URL, error) {
	addr, err := utils.Extract(s.addr, s.listener)
	if err != nil {
		return nil, err
	}

	s.endpoint = &url.URL{
		Scheme: "grpc",
		Host:   addr,
	}
	return s.endpoint, nil
}

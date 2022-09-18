package main

import (
	"context"
	"fmt"

	pb "gin/examples/rpc/test"
	grpcSrv "gin/internal/server"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jonny-mark/gin-micro-mine/pkg/app"
	eagle "github.com/jonny-mark/gin-micro-mine/pkg/app"
	logger "github.com/jonny-mark/gin-micro-mine/pkg/log"
	"google.golang.org/grpc"
	"net/http"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	if in.Name == "" {
		return nil, fmt.Errorf("invalid argument %s", in.Name)
	}
	return &pb.HelloReply{Message: fmt.Sprintf("Hello %+v", in.Name)}, nil
}

func httpServer() error {
	ctxr := context.Background()
	ctx, cancel := context.WithCancel(ctxr)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	//與原生gRPC不同點在這邊，需要做http與grpc的對應
	err := pb.RegisterGreeterHandlerFromEndpoint(ctx, mux, ":9090", opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":7878", mux)
}

func main() {
	cfg := &app.ServerConfig{
		Network:      "tcp",
		Addr:         ":9090",
		ReadTimeout:  200,
		WriteTimeout: 200,
	}
	grpcServer := grpcSrv.NewGRPCServer(cfg)
	srv := &server{}
	pb.RegisterGreeterServer(grpcServer, srv)

	//httpServer := grpcSrv.NewHTTPServer(cfg)
	go httpServer()

	// start app
	app := eagle.New(
		eagle.WithName("eagle"),
		eagle.WithVersion("v1.0.0"),
		eagle.WithLogger(logger.GetLogger()),
		eagle.WithServer(
			grpcServer,
			//httpServer,
		),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}

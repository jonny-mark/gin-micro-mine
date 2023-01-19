package main

import (
	pb "github.com/jonny-mark/gin-micro-mine/api/grpc/user/v1"
	"github.com/jonny-mark/gin-micro-mine/internal/server"
	"github.com/jonny-mark/gin-micro-mine/pkg/app"
	"github.com/jonny-mark/gin-micro-mine/pkg/log"

	"context"
)

type grpcSrv struct {
	pb.UnimplementedUserServiceServer
}

func (se *grpcSrv) LoginByPhone(c context.Context, r *pb.PhoneLoginRequest) (*pb.PhoneLoginReply, error) {
	return &pb.PhoneLoginReply{Ret: string(r.GetPhone()) + r.GetVerifyCode()}, c.Err()
}

func main() {
	cfg := &app.ServerConfig{
		Network:      "tcp",
		Addr:         ":9090",
		ReadTimeout:  200,
		WriteTimeout: 200,
	}
	grpcServer := server.NewGRPCServer(cfg)

	srv := &grpcSrv{}
	pb.RegisterUserServiceServer(grpcServer, srv)
	myApp := app.New(
		app.WithName("test"),
		app.WithVersion("1.0"),
		app.WithLogger(log.GetLogger()),
		app.WithServer(
			// init http server
			//server.NewHTTPServer(&app.Conf.HTTP),

			// init grpc server
			grpcServer,
		),
		//app.WithRegistry(r),
	)
	if err := myApp.Run(); err != nil {
		panic(err)
	}
}

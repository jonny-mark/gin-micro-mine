/**
 * @author jiangshangfang
 * @date 2021/10/24 8:13 PM
 **/
package server

import (
	"gin/pkg/app"
	"gin/pkg/transport/grpc"
)

func NewGRPCServer(cfg *app.ServerConfig) *grpc.Server {
	srv := grpc.NewServer(
		grpc.WithNetwork(cfg.Network),
		grpc.WithAddr(cfg.Addr),
		grpc.WithReadTimeout(cfg.ReadTimeout),
		grpc.WithWriteTimeout(cfg.WriteTimeout),
	)

	//v1.RegisterUserServiceServer(srv, service.Svc.Users().(v1.UserServiceServer))

	return srv
}

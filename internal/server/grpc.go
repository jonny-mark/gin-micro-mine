/**
 * @author jiangshangfang
 * @date 2021/10/24 8:13 PM
 **/
package server

import (
	"gin/pkg/transport/grpc"
	"gin/pkg/app"
)

func NewGRPCServer(cfg *app.ServerConfig) *grpc.Server {
	srv := grpc.NewServer(
		grpc.WithNetwork(cfg.Network),
		grpc.WithAddr(cfg.Addr),
		grpc.WithReadTimeout(cfg.ReadTimeout),
		grpc.WithWriteTimeout(cfg.WriteTimeout),
	)
	return srv
}

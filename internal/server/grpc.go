/**
 * @author jiangshangfang
 * @date 2021/10/24 8:13 PM
 **/
package server

import (
	"gin/pkg/transport/grpc"
	"gin/pkg/log"
	"net"
)

func NewGRPCServer(cfg *grpc.Grpc) *grpc.Server {
	lis, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v1", err)
	}

	grpcServer := grpc.NewServer()
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve grpc server: %v1", err)
	}
	log.Infof("serve grpc server is success, port:%s", cfg.Addr)
	return grpcServer
}

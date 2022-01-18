/**
 * @author jiangshangfang
 * @date 2021/12/30 4:31 PM
 **/
package main

import (
	"context"
	"fmt"
	"log"
	pb "gin/api/grpc/user/v1"

	"google.golang.org/grpc"
	"time"
)

func main() {
	conn, err := grpc.Dial("http://localhost:8085", grpc.WithInsecure(), grpc.WithBlock())

	defer func() {
		_ = conn.Close()
	}()

	client := pb.NewUserServiceClient(conn)

	userReq := &pb.PhoneLoginRequest{
		Phone:      13126963723,
		VerifyCode: 123456,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := client.LoginByPhone(ctx, userReq)
	if err != nil {
		log.Fatalf("[rpc] user login by phone err: %v", err)
	}
	fmt.Printf("UserService LoginByPhone : %+v", reply)
}

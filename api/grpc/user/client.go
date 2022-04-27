/**
 * @author jiangshangfang
 * @date 2021/12/30 4:31 PM
 **/
package main

import (
	"context"
	"fmt"
	pb "gin/api/grpc/user/v1"
	"log"

	"flag"
	"google.golang.org/grpc"
	"time"
)

func main() {
	conn, err := grpc.Dial("http://localhost:8085", grpc.WithInsecure(), grpc.WithBlock())

	defer func() {
		_ = conn.Close()
	}()

	client := pb.NewGreeterClient(conn)

	flag.Parse()
	//拦截器设置
	scsource := `{
		"methodConfig": [{
		  "name": [{"service": "echo.echo","method":"Echo"}],
		  "retryPolicy": {
			  "MaxAttempts": 4,
			  "InitialBackoff": ".01s",
			  "MaxBackoff": ".1s",
			  "BackoffMultiplier": 1.0,
			  "RetryableStatusCodes": [ "UNAVAILABLE" ]
		  }
		}]}`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client.SayHello(ctx)
	conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultServiceConfig(scsource))

	if err != nil {
		fmt.Printf("grpc dial err: %v", err)
		panic("grpc dial err")
	}
	defer func() {
		_ = conn.Close()
	}()

	cli := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.HelloRequest{
		Name: *name,
	}
	reply, err := cli.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("Greeting : %s", reply.GetMessage())

	userReq := &pb.HelloRequest{
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

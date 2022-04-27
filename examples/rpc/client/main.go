package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "gin/examples/rpc/test"
)

const (
	defaultName = "eagle"
)

var (
	addr = flag.String("addr", "localhost:9090", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
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
}

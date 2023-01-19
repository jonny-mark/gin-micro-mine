package main

import (
	"context"
	"fmt"
	pb "github.com/jonny-mark/gin-micro-mine/api/grpc/user/v1"
	"google.golang.org/grpc/credentials/insecure"

	"flag"
	"google.golang.org/grpc"
	"time"
)

func main() {
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
	conn, err := grpc.Dial(
		"http://localhost:901",
		grpc.WithTransportCredentials(insecure.NewCredentials()), //不要在连接中使用SSL/TLS
		grpc.WithBlock(), //等到建立连接（执行同步处理）
		grpc.WithDefaultServiceConfig(scsource))

	defer func() {
		_ = conn.Close()
	}()
	fmt.Print("打印")
	client := pb.NewUserServiceClient(conn)

	flag.Parse()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := client.LoginByPhone(ctx, &pb.PhoneLoginRequest{
		Phone:      1111,
		VerifyCode: "2222",
	})

	if err != nil {
		fmt.Printf("grpc dial err: %v", err)
		panic("grpc dial err")
	}
	fmt.Printf("UserService LoginByPhone : %+v", reply.GetRet())
}

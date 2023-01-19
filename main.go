package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	pb "github.com/jonny-mark/gin-micro-mine/api/grpc/user/v1"
	"github.com/jonny-mark/gin-micro-mine/internal/server"
	"github.com/jonny-mark/gin-micro-mine/pkg/app"
	"github.com/jonny-mark/gin-micro-mine/pkg/config"
	"github.com/jonny-mark/gin-micro-mine/pkg/load/nacos"
	logger "github.com/jonny-mark/gin-micro-mine/pkg/log"
	"github.com/jonny-mark/gin-micro-mine/pkg/redis"
	"github.com/jonny-mark/gin-micro-mine/pkg/storage/orm"
	"github.com/jonny-mark/gin-micro-mine/pkg/trace"
	v "github.com/jonny-mark/gin-micro-mine/pkg/version"
	"github.com/spf13/pflag"
	"log"
	"net/http"
	"os"
)

var (
	cfgDir  = pflag.StringP("config dir", "c", "", "config directory.")
	env     = pflag.StringP("env name", "e", "", "env var name.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	// 把用户传递的命令行参数解析为对应变量的值
	pflag.Parse()
	if *version {
		ver := v.Get()
		marshaled, err := json.MarshalIndent(ver, "", "")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(marshaled))
		return
	}
	// 初始化cfg目录
	config.New(*cfgDir, config.WithEnv(*env))

	// 初始化各组件
	initModule()

	//初始化service
	//service.Svc = service.New(repository.New(orm.GetDB()))
	/**加载资源end**/

	gin.SetMode(app.Conf.Mode)

	//初始化pprof
	go func() {
		fmt.Printf("Listening and serving PProf HTTP on %s\n", app.Conf.PprofPort)
		if err := http.ListenAndServe(app.Conf.PprofPort, http.DefaultServeMux); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen ListenAndServe for PProf, err: %s", err.Error())
		}
	}()

	//client, err := etcdclient.New(etcdclient.Config{
	//	Endpoints: strings.Split(app.Conf.Registry.Endpoints, ","),
	//})
	//if err != nil {
	//	log.Fatalf("etcdclient new failed, err: %s", err.Error())
	//}
	//r := etcd.New(client)

	myApp := app.New(
		app.WithName(app.Conf.Name),
		app.WithVersion(app.Conf.Version),
		app.WithLogger(logger.GetLogger()),
		app.WithServer(
			// init http server
			server.NewHTTPServer(&app.Conf.HTTP),

			// init grpc server
			//grpcServer,
		),
		//app.WithRegistry(r),
	)

	go RunGrpc()
	if err := myApp.Run(); err != nil {
		panic(err)
	}
}

func RunGrpc() {
	grpcServer := server.NewGRPCServer(&app.Conf.GRPC)

	srv := &grpcSrv{}
	pb.RegisterUserServiceServer(grpcServer, srv)
	myApp := app.New(
		app.WithName(app.Conf.Name),
		app.WithVersion(app.Conf.Version),
		app.WithLogger(logger.GetLogger()),
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

type grpcSrv struct {
	pb.UnimplementedUserServiceServer
}

func (se *grpcSrv) LoginByPhone(c context.Context, r *pb.PhoneLoginRequest) (*pb.PhoneLoginReply, error) {
	return &pb.PhoneLoginReply{Ret: fmt.Sprintf("Phone: %+v", r.Phone)}, nil
}

// initModule 初始化组件
func initModule() {
	/**加载资源start**/
	//初始化nacos配置
	nacos.Init()
	//初始化app
	app.Init()
	//初始化日志
	logger.Init()
	//初始化数据库
	orm.Init()
	//初始化RabbitMq
	//rabbitMq.init()
	//初始化redis
	redis.Init()
	//初始化trace
	if app.Conf.EnableTrace {
		trace.Init()
	}
}

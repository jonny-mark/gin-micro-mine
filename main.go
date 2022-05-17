package main

import (
	"encoding/json"
	"fmt"
	"gin/internal/server"
	"gin/pkg/app"
	"gin/pkg/config"
	"gin/pkg/load/nacos"
	logger "gin/pkg/log"
	"gin/pkg/redis"
	"gin/pkg/storage/orm"
	"gin/pkg/trace"
	v "gin/pkg/version"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"log"
	"os"
	"net/http"
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
	spew.Dump(*cfgDir)
	spew.Dump(*env)
	// 初始化cfg目录
	config.New(*cfgDir, config.WithEnv(*env))

	/**加载资源start**/
	//初始化nacos配置
	nacos.Init()
	//初始化app
	app.Init()
	//初始化日志
	logger.Init()
	//初始化数据库
	orm.Init()
	//初始化redis
	redis.Init()
	//初始化trace
	if app.Conf.EnableTrace {
		trace.Init()
	}
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

	// start app
	myApp := app.New(
		app.WithName(app.Conf.Name),
		app.WithVersion(app.Conf.Version),
		app.WithLogger(logger.GetLogger()),
		app.WithServer(
			// init http server
			server.NewHTTPServer(&app.Conf.HTTP),
			//// init grpc server
			//server.NewGRPCServer(&cfg.GRPC),
		),
		//app.WithRegistry(r),
	)

	if err := myApp.Run(); err != nil {
		panic(err)
	}
}

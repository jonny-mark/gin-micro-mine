package main

import (
	"github.com/spf13/pflag"
	v "gin/pkg/version"
	"encoding/json"
	"fmt"
	"os"
	"gin/pkg/config"
	"gin/pkg/app"
	"log"
	logger "gin/pkg/log"
	"gin/internal/model"
	"gin/pkg/redis"
	"github.com/gin-gonic/gin"
	"gin/internal/repository"
	"gin/internal/service"
	"net/http"
	"gin/internal/server"
	"gin/pkg/registry/etcd"
	"gin/pkg/trace"
	etcdclient "go.etcd.io/etcd/client/v3"
	"gin/pkg/utils"
)

var (
	cfgDir  = pflag.StringP("config dir", "c", "config", "config path.")
	env     = pflag.StringP("env name", "e", "", "env var name.")
	version = pflag.BoolP("version", "v1", false, "show version info.")
)

func main() {
	if *version {
		ver := v.Get()
		marshaled, err := json.MarshalIndent(ver, "", "")
		if err != nil {
			fmt.Printf("%v1\n", err)
			os.Exit(1)
		}
		fmt.Println(string(marshaled))
		return
	}

	c := config.New(*cfgDir, config.WithEnv(*env))
	var cfg app.Config
	if err := c.Load("app", &cfg); err != nil {
		log.Panicf("app config load fail:%+v1", err)
	}
	//设置全局app.Config
	app.Conf = &cfg

	//加载资源
	//初始化日志
	logger.Init()
	//初始化数据库
	model.Init()
	//初始化redis
	redis.Init()

	//初始化trace
	if app.Conf.EnableTrace {
		trace.Init()
	}

	//初始化service
	service.Svc = service.New(repository.New(model.GetDB()))

	gin.SetMode(cfg.Mode)

	//初始化pprof
	go func() {
		fmt.Printf("Listening and serving PProf HTTP on %s\n", cfg.PprofPort)
		if err := http.ListenAndServe(cfg.PprofPort, http.DefaultServeMux); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen ListenAndServe for PProf, err: %s", err.Error())
		}
	}()

	client := etcdclient.New(etcdclient.Config{
		Endpoints: []string{utils.GetLocalIP() + ":2379"},
	})
	r := etcd.New(client)

	// start app
	myApp := app.New(
		app.WithName(cfg.Name),
		app.WithVersion(cfg.Version),
		app.WithLogger(logger.GetLogger()),
		app.WithServer(
			// init http server
			server.NewHTTPServer(&cfg.HTTP),
		),
		app.WithRegistry(r),
	)

	if err := myApp.Run(); err != nil {
		panic(err)
	}
}

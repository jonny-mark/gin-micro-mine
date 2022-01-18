package main

import (
	"github.com/spf13/pflag"
	v "gin/pkg/version"
	"encoding/json"
	"fmt"
	"os"
	"gin/pkg/config"
	myApp "gin/pkg/app"
	"log"
	logger "gin/pkg/log"
	"gin/internal/model"
	"gin/pkg/redis"
	"github.com/gin-gonic/gin"
	"gin/internal/repository"
	"gin/internal/service"
	"net/http"
	"gin/internal/server"
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
	logger.Init()
	model.Init()
	redis.Init()
	service.Svc = service.New(repository.New(model.GetDB()))

	gin.SetMode(cfg.Mode)

	// init pprof server
	go func() {
		fmt.Printf("Listening and serving PProf HTTP on %s\n", cfg.PprofPort)
		if err := http.ListenAndServe(cfg.PprofPort, http.DefaultServeMux); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen ListenAndServe for PProf, err: %s", err.Error())
		}
	}()

	// start app
	app := myApp.New(
		eagle.WithName(cfg.Name),
		eagle.WithVersion(cfg.Version),
		eagle.WithLogger(logger.GetLogger()),
		eagle.WithServer(
			// init http server
			server.NewHTTPServer(&cfg.HTTP),
		),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}

/**
 * @author jiangshangfang
 * @date 2021/8/8 4:54 PM
 **/
package core

import (
	"gin/router"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/fvbock/endless"
	"fmt"
	"gin/global"
	"go.uber.org/zap"
)
type server interface {
	ListenAndServe() error
}

func Run()  {
	addr := fmt.Sprintf(":%d",global.Config.System.Addr)
	server := initServer(addr,router.Router())
	global.Log.Info("server run success on ", zap.String("address", addr))
	global.Log.Error(server.ListenAndServe().Error())
}


func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

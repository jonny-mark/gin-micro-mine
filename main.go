package main

import (
	"gin/global"
	"gin/pkg/viper"
	"gin/pkg/logger"
	"gin/pkg/core"
)

func main()  {
	global.Vip = viper.Viper()
	global.Log = logger.Zap()

    core.Run()
}

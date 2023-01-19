package sentinel

import (
	sentinelConfig "github.com/alibaba/sentinel-golang/core/config"
	//sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/jonny-mark/gin-micro-mine/pkg/config"
	"log"
)

func Init() {
	var c sentinelConfig.Entity
	if err := config.Load("sentinal", &c); err != nil {
		log.Panicf("redis config load %+v", err)
	}
	//sentinel.InitWithConfigFile()

	//err := sentinel.InitWithConfig()
	//if err != nil {
	//	// 初始化 Sentinel 失败
	//}
}

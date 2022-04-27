/**
 * @author jiangshangfang
 * @date 2022/3/3 4:41 PM
 **/
package sentinel

import (
	//sentinel "github.com/alibaba/sentinel-golang/api"
	"gin/pkg/config"
	sentinelConfig "github.com/alibaba/sentinel-golang/core/config"
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

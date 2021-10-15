/**
 * @author jiangshangfang
 * @date 2021/7/29 11:58 AM
 **/
package global

import (
	"gin/global/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Config config.Server
	Vip   *viper.Viper
	Log   *zap.Logger
)

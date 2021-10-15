/**
 * @author jiangshangfang
 * @date 2021/7/29 2:31 PM
 **/
package viper

import (
	"github.com/spf13/viper"
	"gin/common/constant"
	"fmt"
	"flag"
	"github.com/fsnotify/fsnotify"
	"gin/global"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file")
		flag.Parse()
		if config == "" {
			config = constant.ConfigPath + constant.ConfigFile
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType(constant.ConfigType)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("config file changed: %s \n", in.Name)
		if err := v.Unmarshal(&global.Config); err != nil {
			fmt.Printf("config unmarshal: %s \n", err)
		}
	})
	if err := v.Unmarshal(&global.Config); err != nil {
		fmt.Printf("config unmarshal: %s \n", err)
	}
	return v
}

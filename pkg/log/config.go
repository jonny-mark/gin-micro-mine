/**
 * @author jiangshangfang
 * @date 2021/10/17 9:14 PM
 **/
package log

import (
	"github.com/jonny-mark/gin-micro-mine/internal/constant"
	"github.com/jonny-mark/gin-micro-mine/pkg/config"
	"github.com/jonny-mark/gin-micro-mine/pkg/load/nacos"
	"gopkg.in/yaml.v3"
	"log"
)

type Config struct {
	Name             string `yaml:"Name"`
	Development      bool   `yaml:"Development"`
	Level            string `yaml:"Level"`
	Format           string `yaml:"Format"`
	Stacktrace       bool   `yaml:"Stacktrace"`
	LinkName         string `yaml:"LinkName"`
	Prefix           string `yaml:"Prefix"`
	Director         string `yaml:"Director"`
	LogRollingPolicy string `yaml:"LogRollingPolicy"`
	LoggerInfoFile   string `yaml:"LoggerInfoFile"`
	LoggerWarnFile   string `yaml:"LoggerWarnFile"`
	LoggerErrorFile  string `yaml:"LoggerErrorFile"`
	MaxAge           int64  `yaml:"MaxAge"`
}

// 初始化
func Init() Logger {
	var cfg Config
	if nacos.NacosClient.Enable {
		context, err := nacos.NacosClient.LoadConfiguration(constant.NacosLoggerKey)
		if err != nil {
			log.Panicf("load logger conf err: %v", err)
		}
		if err := yaml.Unmarshal([]byte(context), &cfg); err != nil {
			log.Panicf("load logger conf unmarshal err: %v", err)
		}
		listenConfiguration(constant.NacosLoggerKey, &cfg)
	} else {
		if err := config.Load(constant.LoggerKey, &cfg); err != nil {
			log.Panicf("load logger conf err: %v", err)
		}
	}
	logger = newZapLogger(&cfg)
	return logger
}

func listenConfiguration(name string, cfg *Config) {
	context, err := nacos.NacosClient.ListenConfiguration(name)
	if err != nil {
		log.Panicf("load logger conf err: %v", err)
	}
	if err := yaml.Unmarshal([]byte(context), cfg); err != nil {
		log.Panicf("load logger conf unmarshal err: %v", err)
	}
}

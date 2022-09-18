/**
 * @author jiangshangfang
 * @date 2021/10/21 4:35 PM
 **/
package orm

import (
	"gin/internal/constant"
	"gin-micro-mine/pkg/config"
	"gin-micro-mine/pkg/load/nacos"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"log"
	"time"
)

type Config struct {
	Name            string        `yaml:"Name"`
	Addr            string        `yaml:"Addr"`
	UserName        string        `yaml:"UserName"`
	Password        string        `yaml:"Password"`
	MaxIdleConn     int           `yaml:"MaxIdleConn"`
	ConnMaxLifeTime time.Duration `yaml:"ConnMaxLifeTime"`
	MaxOpenConn     int           `yaml:"MaxOpenConn"`
	ShowLog         bool          `yaml:"ShowLog"`
	SlowThreshold   int64         `yaml:"SlowThreshold"`
}

// Init 初始化数据库
func Init() *gorm.DB {
	var cfg Config
	if nacos.NacosClient.Enable {
		context, err := nacos.NacosClient.LoadConfiguration(constant.NacosDatabaseKey)
		if err != nil {
			log.Panicf("load database conf err: %v", err)
		}
		if err := yaml.Unmarshal([]byte(context), &cfg); err != nil {
			log.Panicf("load database conf unmarshal err: %v", err)
		}
		listenConfiguration(constant.NacosDatabaseKey, &cfg)
	} else {
		if err := config.Load(constant.DatabaseKey, &cfg); err != nil {
			log.Panicf("database config load %+v", err)
		}
	}
	DB = NewMysql(&cfg)
	return DB
}

//监听nacos的变化
func listenConfiguration(name string, cfg *Config) {
	context, err := nacos.NacosClient.ListenConfiguration(name)
	if err != nil {
		log.Panicf("load database conf err: %v", err)
	}
	if err := yaml.Unmarshal([]byte(context), cfg); err != nil {
		log.Panicf("load database conf unmarshal err: %v", err)
	}
}

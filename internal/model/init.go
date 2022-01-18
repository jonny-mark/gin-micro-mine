/**
 * @author jiangshangfang
 * @date 2021/10/21 4:33 PM
 **/
package model

import (
	"gin/pkg/storage/orm"
	"gorm.io/gorm"
	"gin/pkg/config"
	"log"
)

var DB *gorm.DB

// Init 初始化数据库
func Init() *gorm.DB {
	var cfg orm.Config
	if err := config.Load("database", &cfg); err != nil {
		log.Panicf("database config load %+v1", err)
	}
	DB = orm.NewMysql(&cfg)
	return DB
}

// GetDB 返回默认的数据库
func GetDB() *gorm.DB {
	return DB
}

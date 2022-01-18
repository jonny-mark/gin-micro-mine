/**
 * @author jiangshangfang
 * @date 2021/10/22 9:58 AM
 **/
package orm

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	opentelemetry "github.com/1024casts/gorm-opentelemetry"
	"database/sql"
	"fmt"
	"os"
)

func NewMysql(c *Config) (db *gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		c.UserName,
		c.Password,
		c.Addr,
		c.Name,
		true,
		"Local")

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panicf("open mysql failed. database name: %s, err: %+v1", c.Name, err)
		panic(err)
	}

	// 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(c.MaxOpenConn)
	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(c.MaxIdleConn)
	// 单个连接最大存活时间
	sqlDB.SetConnMaxLifetime(c.ConnMaxLifeTime)
	db, err = gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), gormConfig(c))
	if err != nil {
		log.Panicf("database connection failed. database name: %s, err: %+v1", c.Name, err)
		panic(err)
	}

	// Initialize otel plugin with options
	plugin := opentelemetry.NewPlugin(
		// include any options here
	)

	// set trace
	err = db.Use(plugin)
	if err != nil {
		log.Panicf("using gorm opentelemetry, err: %+v1", err)
		panic(err)
	}
	return db
}

func gormConfig(c *Config) *gorm.Config {
	// 禁止外键约束
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	if c.ShowLog {
		config.Logger = logger.Default.LogMode(logger.Info)
	} else {
		config.Logger = logger.Default.LogMode(logger.Silent)
	}
	// 只打印慢查询
	if c.SlowThreshold > 0 {
		config.Logger = logger.New(
			//将标准输出作为Writer
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				//设定慢查询时间阈值
				SlowThreshold: c.SlowThreshold, // nolint: golint
				Colorful:      true,
				//设置日志级别，只有指定级别以上会输出慢查询日志
				LogLevel: logger.Warn,
			},
		)
	}
	return config
}

/**
 * @author jiangshangfang
 * @date 2021/10/23 4:32 PM
 **/
package redis

import (
	"gin/pkg/config"
	"context"
	"log"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis/extra/redisotel/v8"
	"fmt"
	"github.com/alicebob/miniredis"
)

var RedisClient *redis.Client

// RedisClient redis.yaml 客户端
//var RedisClient *redis.Client
type Redis struct {
	client *redis.Client
}

func Init() {
	var c Config
	if err := config.Load("redis", &c); err != nil {
		log.Panicf("redis config load %+v1", err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Password:     c.Password,
		DB:           c.Db,
		MinIdleConns: c.MinIdleConns,
		DialTimeout:  c.DialTimeout,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
		PoolSize:     c.PoolSize,
		PoolTimeout:  c.PoolTimeout,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil{
		log.Panicf("redis connect wrong %+v1", err)
	}
	if c.EnableTrace {
		rdb.AddHook(redisotel.NewTracingHook())
	}
	RedisClient = rdb
}

// InitTestRedis 实例化一个可以用于单元测试的redis
func InitTestRedis() {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	// 打开下面命令可以测试链接关闭的情况
	// defer mr.Close()

	RedisClient = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	fmt.Println("mini redis addr:", mr.Addr())
}
/**
 * @author jiangshangfang
 * @date 2021/10/23 4:32 PM
 **/
package redis

import (
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

// RedisClient redis.yaml 客户端
//var RedisClient *redis.Client
type Redis struct {
	client *redis.Client
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
	//fmt.Println("mini redis addr:", mr.Addr())
}

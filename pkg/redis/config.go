package redis

import (
	"context"
	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
	"github.com/jonny-mark/gin-micro-mine/internal/constant"
	"github.com/jonny-mark/gin-micro-mine/pkg/config"
	"github.com/jonny-mark/gin-micro-mine/pkg/load/nacos"
	"gopkg.in/yaml.v3"
	"log"
	"time"
)

type Config struct {
	Addr         string        `yaml:"Addr"`
	Password     string        `yaml:"Password"`
	Db           int           `yaml:"Db"`
	MaxRetries   int           `yaml:"MaxRetries"`
	PoolSize     int           `yaml:"PoolSize"`
	PoolTimeout  time.Duration `yaml:"PoolTimeout"`
	MinIdleConns int           `yaml:"MinIdleConns"`
	DialTimeout  time.Duration `yaml:"DialTimeout"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
	EnableTrace  bool          `yaml:"EnableTrace"`
}

func Init() {
	var c Config
	if nacos.NacosClient.Enable {
		ctx, err := nacos.NacosClient.LoadConfiguration(constant.NacosRedisKey)
		if err != nil {
			log.Panicf("load redis conf err: %v", err)
		}
		if err := yaml.Unmarshal([]byte(ctx), &c); err != nil {
			log.Panicf("load redis conf unmarshal err: %v", err)
		}

		listenConfiguration(constant.NacosRedisKey, &c)
	} else {
		if err := config.Load(constant.RedisKey, &c); err != nil {
			log.Panicf("redis config load %+v", err)
		}
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
	if err != nil {
		log.Panicf("redis connect wrong %+v", err)
	}
	if c.EnableTrace {
		rdb.AddHook(redisotel.NewTracingHook())
	}
	RedisClient = rdb
}

// 监听nacos的变化
func listenConfiguration(name string, cfg *Config) {
	ctx, err := nacos.NacosClient.ListenConfiguration(name)
	if err != nil {
		log.Panicf("load redis conf err: %v", err)
	}
	if err := yaml.Unmarshal([]byte(ctx), cfg); err != nil {
		log.Panicf("load redis conf unmarshal err: %v", err)
	}
}

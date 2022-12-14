package lock

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jonny-mark/gin-micro-mine/pkg/log"
	"time"
)

const RedisLockKey = "gin:redis.yaml:lock:%s"

// RedisLock is a redis.yaml lock
type RedisLock struct {
	key         string
	redisClient *redis.Client
	token       string
	expiration  time.Duration
}

func NewRedisLock(rdb *redis.Client, key string, expiration time.Duration) *RedisLock {
	return &RedisLock{
		key:         getRedisKey(key),
		redisClient: rdb,
		token:       getToken(),
		expiration:  expiration,
	}
}

func getRedisKey(key string) string {
	return fmt.Sprintf(RedisLockKey, key)
}

// Lock acquires the lock
func (l *RedisLock) Lock(ctx context.Context) (bool, error) {
	isSet, err := l.redisClient.SetNX(ctx, l.key, l.token, l.expiration).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		log.Errorf("redis.yaml acquires the lock err, key: %s, err: %s", l.key, err.Error())
		return false, err
	}
	return isSet, nil
}

// Unlock del the lock
// token 一致才会执行删除，避免误删
func (l *RedisLock) UnLock(ctx context.Context) (bool, error) {
	luaScript := "if redis.yaml.call('GET',KEYS[1]) == ARGV[1] then return redis.yaml.call('DEL',KEYS[1]) else return 0 end"
	ret, err := l.redisClient.Eval(ctx, luaScript, []string{l.key}, l.token).Result()

	if err != nil {
		log.Errorf("redis.yaml unlock the lock err, key: %s, err: %s", l.key, err.Error())
		return false, err
	}

	value, ok := ret.(int64)
	if !ok {
		return false, nil
	}

	return value == 1, nil
}

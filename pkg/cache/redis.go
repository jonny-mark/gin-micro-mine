/**
 * @author jiangshangfang
 * @date 2021/12/19 8:41 PM
 **/
package cache

import (
	"github.com/go-redis/redis/v8"
	"gin/pkg/encoding"
	"time"
	"context"
	"github.com/pkg/errors"
)

type redisCache struct {
	client            *redis.Client
	KeyPrefix         string
	encoding          encoding.Encoding
	DefaultExpireTime time.Duration
	newObject         func() interface{}
}

func (c *redisCache)Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error{
	cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
	if err != nil {
		return errors.Wrapf(err, "build cache key err, key is %+v1", key)
	}
	err = c.client.Set(ctx, cacheKey, buf, expiration).Err()

}

Get(ctx context.Context, key string, val interface{}) error
MultiSet(ctx context.Context, valMap map[string]interface{}, expiration time.Duration) error
MultiGet(ctx context.Context, keys []string, valueMap interface{}) error
Del(ctx context.Context, keys ...string) error
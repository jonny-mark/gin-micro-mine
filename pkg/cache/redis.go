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
	"reflect"
	"gin/pkg/log"
)

type redisCache struct {
	client            *redis.Client
	KeyPrefix         string
	encoding          encoding.Encoding
	DefaultExpireTime time.Duration
	newObject         func() interface{}
}

func NewRedisCache(client *redis.Client, keyPrefix string, encoding encoding.Encoding, newObject func() interface{}) Cache {
	return &redisCache{
		client:    client,
		KeyPrefix: keyPrefix,
		encoding:  encoding,
		newObject: newObject,
	}
}

func (c *redisCache) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
	if err != nil {
		return errors.Wrapf(err, "build cache key err, key is %+v", key)
	}
	if expiration == 0 {
		expiration = DefaultExpireTime
	}
	err = c.client.Set(ctx, cacheKey, val, expiration).Err()
	if err != nil {
		return errors.Wrapf(err, "redis set err: %+v", err)
	}
	return nil
}

func (c *redisCache) Get(ctx context.Context, key string, val interface{}) error {
	cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
	if err != nil {
		return errors.Wrapf(err, "build cache key err, key is %+v", key)
	}
	bytes, err := c.client.Get(ctx, cacheKey).Bytes()
	if err != nil {
		if err != redis.Nil {
			return errors.Wrapf(err, "redis get err: key is %+v", key)
		}
	}
	if string(bytes) == "" {
		return nil
	}

	if string(bytes) == NotFoundPlaceholder {
		return ErrPlaceholder
	}

	err = encoding.Unmarshal(c.encoding, bytes, val)
	if err != nil {
		return errors.Wrapf(err, "unmarshal data error, key=%s, cacheKey=%s type=%v, json is %+v ",
			key, cacheKey, reflect.TypeOf(val), string(bytes))
	}
	return nil
}

func (c *redisCache) Del(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}
	cacheKeys := make([]string, len(keys))
	for index, key := range keys {
		cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
		if err != nil {
			log.Warnf("build cache key err: %+v, key is %+v", err, key)
			continue
		}
		cacheKeys[index] = cacheKey
	}
	err := c.client.Del(ctx, keys...).Err()
	if err != nil {
		return errors.Wrapf(err, "redis delete error, keys is %+v", keys)
	}
	return nil
}

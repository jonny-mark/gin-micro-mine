/**
 * @author jiangshangfang
 * @date 2022/5/22 17:33
 **/
package cache

import (
	"fmt"
	"gin/internal/model/user"
	"gin/pkg/cache"
	"go.opentelemetry.io/otel/trace"
	"gin/pkg/encoding"
	"go.opentelemetry.io/otel"
	"gin/pkg/redis"
	"time"
	"context"
	"gin/pkg/log"
	"gin/internal/constant"
)

// Cache cache
type Cache struct {
	cache  cache.Cache
	tracer trace.Tracer
}

// NewUserCache new一个用户cache
func NewUserCache() *Cache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	return &Cache{
		cache: cache.NewRedisCache(redis.RedisClient, cachePrefix, jsonEncoding, func() interface{} {
			return &user.UsersModel{}
		},
		),
		tracer: otel.Tracer("user cache"),
	}
}

// GetUserBaseCacheKey 获取cache key
func (c *Cache) GetUserBaseCacheKey(userID uint64) string {
	return fmt.Sprintf(constant.PrefixUserBaseCacheKey, userID)
}

// SetUserBaseCache 写入用户cache
func (c *Cache) SetUserBaseCache(ctx context.Context, userID uint64, user *user.UsersModel, duration time.Duration) error {
	ctx, span := c.tracer.Start(ctx, "SetUserBaseCache")
	defer span.End()

	if user == nil || user.ID == 0 {
		return nil
	}
	cacheKey := c.GetUserBaseCacheKey(userID)
	err := c.cache.Set(ctx, cacheKey, &user, duration)
	if err != nil {
		return err
	}
	return nil
}

// GetUserBaseCache 获取用户cache
func (c *Cache) GetUserBaseCache(ctx context.Context, userID uint64) (data *user.UsersModel, err error) {
	ctx, span := c.tracer.Start(ctx, "GetUserBaseCache")
	defer span.End()

	cacheKey := c.GetUserBaseCacheKey(userID)
	err = c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		log.WithContext(ctx).Warnf("get err from redis, err: %+v", err)
		return nil, err
	}
	return data, nil
}


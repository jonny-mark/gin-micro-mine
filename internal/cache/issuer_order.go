/**
 * @author jiangshangfang
 * @date 2022/2/21 5:00 PM
 **/
package cache

import (
	"context"
	"fmt"
	"gin/internal/model"
	"gin/pkg/cache"
	"gin/pkg/encoding"
	"gin/pkg/log"
	"gin/pkg/redis"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const PrefixIssuerOrderBaseCacheKey = "IssuerOrder:%s"

// Cache cache
type Cache struct {
	cache  cache.Cache
	tracer trace.Tracer
}

func NewIssuerOrderCache() *Cache {
	cachePrefix := ""
	return &Cache{
		cache: cache.NewRedisCache(redis.RedisClient, cachePrefix, encoding.JSONEncoding{}, func() interface{} {
			return &model.IssuerOrderModel{}
		}),
		tracer: otel.Tracer("IssuerOrder cache"),
	}
}

func (c *Cache) GetIssuerOrderCacheKey(OrderSn string) string {
	return fmt.Sprintf(PrefixIssuerOrderBaseCacheKey, OrderSn)
}

func (c *Cache) GetIssuerOrderCache(ctx context.Context, orderSn string) (data *model.IssuerOrderModel, err error) {
	ctx, span := c.tracer.Start(ctx, "GetIssuerOrderCache")
	defer span.End()

	cacheKey := c.GetIssuerOrderCacheKey(orderSn)
	err = c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		log.WithContext(ctx).Warnf("get err from redis, err: %+v", err)
		return nil, err
	}
	return
}

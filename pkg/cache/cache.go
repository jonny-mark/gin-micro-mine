/**
 * @author jiangshangfang
 * @date 2021/12/19 6:38 PM
 **/
package cache

import (
	"context"
	"time"
)

var (
	// DefaultClient 生成一个缓存客户端，其中keyPrefix 一般为业务前缀
	DefaultClient Cache
)

//定义cache驱动接口
type Cache interface {
	Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, val interface{}) error
	MultiSet(ctx context.Context, valMap map[string]interface{}, expiration time.Duration) error
	MultiGet(ctx context.Context, keys []string, valueMap interface{}) error
	Del(ctx context.Context, keys ...string) error
}

func Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error{
	return DefaultClient.Set(ctx,key,val,expiration)
}

func Get(ctx context.Context, key string, val interface{}) error {
	return DefaultClient.Get(ctx, key, val)
}

func MultiSet(ctx context.Context, valMap map[string]interface{}, expiration time.Duration) error {
	return DefaultClient.MultiSet(ctx, valMap, expiration)
}

func MultiGet(ctx context.Context, keys []string, valueMap interface{}) error {
	return DefaultClient.MultiGet(ctx, keys, valueMap)
}

func Del(ctx context.Context, keys ...string) error {
	return DefaultClient.Del(ctx, keys...)
}
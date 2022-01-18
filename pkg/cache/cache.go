/**
 * @author jiangshangfang
 * @date 2021/12/19 6:38 PM
 **/
package cache

import (
	"context"
	"time"
	"errors"
)

var (
	// DefaultClient 生成一个缓存客户端
	DefaultClient Cache

	// DefaultExpireTime 默认过期时间
	DefaultExpireTime = time.Hour * 24

	//展位标识
	NotFoundPlaceholder = "*"
	ErrPlaceholder = errors.New("cache: placeholder")
)

//定义cache驱动接口
type Cache interface {
	Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, val interface{}) error
	Del(ctx context.Context, keys ...string) error
}

func Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error{
	return DefaultClient.Set(ctx,key,val,expiration)
}

func Get(ctx context.Context, key string, val interface{}) error {
	return DefaultClient.Get(ctx, key, val)
}

func Del(ctx context.Context, keys ...string) error {
	return DefaultClient.Del(ctx, keys...)
}
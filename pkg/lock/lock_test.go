/**
 * @author jiangshangfang
 * @date 2021/12/1 5:04 PM
 **/
package lock

import (
	"testing"
	"gin/pkg/redis"
	"context"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/davecgh/go-spew/spew"
)

func TestLockWithDefaultTimeout(t *testing.T) {
	redis.InitTestRedis()

	lock := NewRedisLock(redis.RedisClient, "33333")

	ok, err := lock.Lock(context.Background(), 5*time.Second)

	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Fatal("lock is not ok")
	}
	t.Logf("lock is %t", ok)

	ok, err = lock.UnLock(context.Background())
	if err != nil {
		t.Error(err)
	}
	spew.Dump(lock)
	if !ok {
		t.Fatal("Unlock is not ok")
	}
	t.Logf("Unlock is %t", ok)
}

func TestLockWithTimeout(t *testing.T) {
	redis.InitTestRedis()

	t.Run("should lock/unlock success", func(t *testing.T) {
		ctx := context.Background()
		lock1 := NewRedisLock(redis.RedisClient, "4444")
		ok, err := lock1.Lock(ctx, 5*time.Second)
		assert.Nil(t, err)
		assert.True(t, ok)

		ok, err = lock1.UnLock(ctx)
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("should unlock failed", func(t *testing.T) {
		ctx := context.Background()
		lock2 := NewRedisLock(redis.RedisClient, "5555")
		ok, err := lock2.Lock(ctx, 2*time.Second)
		assert.Nil(t, err)
		assert.True(t, ok)

		time.Sleep(3*time.Second)

		ok, err = lock2.UnLock(ctx)
		t.Logf("Unlock is %t", ok)

		assert.Nil(t, err)
		assert.False(t, ok)
	})
}

/**
 * @author jiangshangfang
 * @date 2021/12/1 5:58 PM
 **/
package lock

import (
	"context"
	"fmt"
	"github.com/jonny-mark/gin-micro-mine/pkg/log"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"time"
)

const EtcdLockKey = "gin:etcd:lock:%s"

// EtcdLock define a etcd lock
type EtcdLock struct {
	sess *concurrency.Session
	mu   *concurrency.Mutex
}

func NewEtcdLock(client *clientv3.Client, key string, opts ...concurrency.SessionOption) (mutex *EtcdLock, err error) {
	mutex = &EtcdLock{}

	mutex.sess, err = concurrency.NewSession(client, opts...)
	if err != nil {
		return
	}
	mutex.mu = concurrency.NewMutex(mutex.sess, getEtcdKey(key))
	return
}

// Lock acquires the lock
func (l *EtcdLock) Lock(ctx context.Context, timeout time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return true, l.mu.Lock(ctx)
}

// Unlock del the lock
func (l *EtcdLock) UnLock(ctx context.Context) (bool, error) {
	err := l.mu.Unlock(ctx)
	if err != nil {
		log.Errorf("etcd unlock the lock err, err: %s", err.Error())
		return false, err
	}
	return true, l.sess.Close()
}

func getEtcdKey(key string) string {
	return fmt.Sprintf(EtcdLockKey, key)
}

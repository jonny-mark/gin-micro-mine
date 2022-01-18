/**
 * @author jiangshangfang
 * @date 2021/12/1 3:19 PM
 **/
package lock

import (
	"context"
	"time"
)

//Lock define common func
type Lock interface {
	Lock(ctx context.Context, timeout time.Duration) (bool, error)
	UnLock(ctx context.Context) (bool, error)
}

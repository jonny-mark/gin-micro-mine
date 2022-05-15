/**
 * @author jiangshangfang
 * @date 2022/2/12 5:45 PM
 **/
package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"github.com/vearne/gin-timeout"
)

// Timeout 超时中间件
func Timeout(t time.Duration) gin.HandlerFunc {
	// see:
	// https://github.com/vearne/gin-timeout
	// https://vearne.cc/archives/39135
	// https://github.com/gin-contrib/timeout
	return timeout.Timeout(
		timeout.WithTimeout(t),
	)
}

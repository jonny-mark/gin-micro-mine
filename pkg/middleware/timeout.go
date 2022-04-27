/**
 * @author jiangshangfang
 * @date 2022/2/12 5:45 PM
 **/
package middleware

import (
	"context"
	"gin/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Timeout 超时中间件
func Timeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		defer func() {
			if ctx.Err() == context.DeadlineExceeded {
				// 标记响应码 中断请求
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.AbortWithStatusJSON(http.StatusGatewayTimeout, app.Response{
					Code:    http.StatusGatewayTimeout,
					Message: ctx.Err().Error(),
				})
			}
		}()

		// 包装后的ctx替换
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

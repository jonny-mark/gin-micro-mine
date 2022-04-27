/**
 * @author jiangshangfang
 * @date 2022/4/16 10:46 PM
 **/
package middleware

import (
	"github.com/gin-gonic/gin"
	"gin/pkg/app"
	"gin/pkg/errcode"
)

// 登陆鉴权 token认证
func Access() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 必须登录
		token := c.GetHeader("X-APP-TOKEN")
		if token == "" {
			token = c.GetHeader("x-app-token")
		}

		if token == "" {
			app.NewResponse().Error(c, errcode.InvalidTokenError)
			c.Abort()
			return
		}


		c.Next()
	}
}

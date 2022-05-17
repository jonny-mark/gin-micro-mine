/**
 * @author jiangshangfang
 * @date 2022/4/16 10:46 PM
 **/
package middleware

import (
	"github.com/gin-gonic/gin"
	"gin/pkg/app"
	"gin/pkg/errcode"
	"gin/internal/repository/user"
	"errors"
	"gorm.io/gorm"
)

// 登陆鉴权 token认证
func TokenAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 必须登录
		token := c.GetHeader("X-APP-TOKEN")
		if token == "" {
			token = c.GetHeader("x-app-token")
		}

		if token == "" {
			app.NewResponse().Error(c, errcode.EmptyTokenError)
			c.Abort()
			return
		}

		userInfo, err := user.FindValidOneByToken(token)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			app.NewResponse().Error(c, errcode.InvalidTokenError)
			c.Abort()
			return
		}
		if err != nil {
			app.NewResponse().Error(c, err)
			c.Abort()
			return
		}

		if userInfo == nil {
			app.NewResponse().Error(c, errcode.InvalidTokenError)
			c.Abort()
			return
		}

		c.Next()
	}
}

// sign签名认证
func SignAccess() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}


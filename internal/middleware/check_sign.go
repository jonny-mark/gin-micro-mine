/**
 * @author jiangshangfang
 * @date 2022/4/27 5:26 PM
 **/
package middleware

import (
	"github.com/gin-gonic/gin"
	"gin/pkg/app"
	"gin/pkg/errcode"
	"gin/internal/repository/user"
	"gorm.io/gorm"
)

// sign签名认证
func Sign() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 必须登录
		appKey := c.GetHeader("X-APP-KEY")
		if appKey == "" {
			appKey = c.GetHeader("x-app-key")
		}

		if appKey == "" {
			app.NewResponse().Error(c, errcode.)
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

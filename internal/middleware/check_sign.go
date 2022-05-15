/**
 * @author jiangshangfang
 * @date 2022/4/27 5:26 PM
 **/
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/davecgh/go-spew/spew"
)

// sign签名认证
func Sign() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取所有参数
		body := c.Request.GetBody
		spew.Dump(body())
		//if appKey == "" {
		//	appKey = c.GetHeader("x-app-key")
		//}
		//
		//if appKey == "" {
		//	app.NewResponse().Error(c, errcode.InvalidTokenError)
		//	c.Abort()
		//	return
		//}
		//
		//userInfo, err := user.FindValidOneByToken(token)
		//if errors.Is(err, gorm.ErrRecordNotFound) {
		//	app.NewResponse().Error(c, errcode.InvalidTokenError)
		//	c.Abort()
		//	return
		//}
		//if err != nil {
		//	app.NewResponse().Error(c, err)
		//	c.Abort()
		//	return
		//}
		//
		//if userInfo == nil {
		//	app.NewResponse().Error(c, errcode.InvalidTokenError)
		//	c.Abort()
		//	return
		//}

		c.Next()
	}
}

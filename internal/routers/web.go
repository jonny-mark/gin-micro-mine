/**
 * @author jiangshangfang
 * @date 2021/8/8 5:20 PM
 **/
package routers

import (
	"gin/internal/middleware"
	pubMiddleware "gin/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gin/internal/web/device"
	"gin/internal/web"
)

func LoadWebRouter(g *gin.Engine) {
	g.GET("/", web.Index)

	apiRouter := g.Group("api/")
	apiRouter.Use(middleware.Translations())

	// 用户登录认证
	apiRouter.Use(pubMiddleware.JWTAuth())
	{

	}

	innerRouter := g.Group("inner/")
	// 内部服务签名认证
	innerRouter.Use(middleware.SignAccess())
	{

	}

	mpRouter := g.Group("mp/")
	// 小程序token认证
	mpRouter.Use(middleware.TokenAccess())
	{
		mpRouter.POST("/device-check/create", device.Create)
	}
}

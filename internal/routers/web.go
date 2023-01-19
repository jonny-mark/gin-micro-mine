package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jonny-mark/gin-micro-mine/internal/middleware"
	"github.com/jonny-mark/gin-micro-mine/internal/web"
	"github.com/jonny-mark/gin-micro-mine/internal/web/device"
	pubMiddleware "github.com/jonny-mark/gin-micro-mine/pkg/middleware"
)

func LoadWebRouter(g *gin.Engine) {
	g.GET("/", web.Index)

	apiRouter := g.Group("api/")
	apiRouter.Use(middleware.Translations())

	// 用户登录认证
	apiRouter.Use(pubMiddleware.JWTAuth())
	{
		//apiRouter.POST("")
	}

	innerRouter := g.Group("inner/")
	// 内部服务签名认证
	innerRouter.Use(middleware.SignAccess())
	{
		innerRouter.POST("/signTest", web.SignText)
	}

	mpRouter := g.Group("mp/")
	// 小程序token认证
	mpRouter.Use(middleware.TokenAccess())
	{
		mpRouter.POST("/device-check/create", device.Create)
	}
}

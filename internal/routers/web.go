/**
 * @author jiangshangfang
 * @date 2021/8/8 5:20 PM
 **/
package routers

import (
	"gin/internal/middleware"
	"gin/internal/web"
	"github.com/gin-gonic/gin"
	"gin/internal/web/device"
)

func LoadWebRouter(g *gin.Engine) *gin.Engine {
	router := g
	router.Use(middleware.Translations())

	// 静态资源，主要是图片
	//router.Use(static.Serve("/static", static.LocalFile("./static", false)))

	router.GET("/", web.Index)

	router.POST("/device-check/create", device.Create)
	return router
}

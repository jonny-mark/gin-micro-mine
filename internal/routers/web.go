/**
 * @author jiangshangfang
 * @date 2021/8/8 5:20 PM
 **/
package routers

import (
	"github.com/gin-gonic/gin"
	"gin/internal/web"
	"github.com/gin-contrib/static"
)

func LoadWebRouter(g *gin.Engine) *gin.Engine {
	router := g
	//// 404
	//router.NoRoute(func(c *gin.Context) {
	//	web.Error404(c)
	//})
	//router.NoMethod(func(c *gin.Context) {
	//	web.Error404(c)
	//})

	// 静态资源，主要是图片
	router.Use(static.Serve("/static", static.LocalFile("./static", false)))

	router.GET("/",web.Index)
	return router
}

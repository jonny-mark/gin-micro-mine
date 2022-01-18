/**
 * @author jiangshangfang
 * @date 2021/8/8 4:53 PM
 **/
package router

import (
	"github.com/gin-gonic/gin"
	"gin/common"
	//mw "gin/internal/middleware"
	"gin/pkg/middleware"
)

func NewRouter() *gin.Engine {
	g := gin.New()
	// 使用中间件
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middleware.Logging)


	//g.Use(mw.Tracing(&config.Conf.Trace))


	// 404 Handler.
	//g.NoRoute(api.RouteNotFound)
	//g.NoMethod(api.RouteNotFound)

	// 静态资源，主要是图片
	g.Static("/static", "./static")
	// HealthCheck 健康检查路由
	g.GET("/health", common.HealthCheck)

	return g
}

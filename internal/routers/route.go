/**
 * @author jiangshangfang
 * @date 2021/8/8 4:53 PM
 **/
package routers

import (
	"github.com/gin-gonic/gin"
	"gin/common"
	mw "gin/internal/middleware"
	"gin/pkg/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gin/pkg/app"
)

func NewRouter() *gin.Engine {
	g := gin.New()
	// 使用中间件
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middleware.Logging())
	g.Use(middleware.RequestId())
	g.Use(middleware.Tracing(app.Conf.Name))
	g.Use(mw.Translations())

	// 加载web路由
	LoadWebRouter(g)
	// 404 Handler.
	g.NoRoute(common.RouteNotFound)
	g.NoMethod(common.RouteNotFound)

	// HealthCheck 健康检查路由
	g.GET("/health", common.HealthCheck)
	// 通过 grafana 可视化查看 prometheus 的监控数据，使用插件6671查看
	g.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return g
}

/**
 * @author jiangshangfang
 * @date 2021/8/8 4:53 PM
 **/
package routers

import (
	"gin/pkg/app"
	"gin/pkg/middleware"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"time"
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
	g.Use(middleware.Metrics(app.Conf.Name))
	g.Use(middleware.Timeout(3 * time.Second))
	g.Use(middleware.JWTAuth())
	g.Use(middleware.SignAuth())

	// 加载web路由
	LoadWebRouter(g)
	// 404 Handler.
	g.NoRoute(app.RouteNotFound)
	g.NoMethod(app.RouteNotFound)

	// swagger api docs
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// pprof router 性能分析路由
	// 默认关闭，开发环境下可以打开
	// 访问方式: HOST/debug/pprof
	// 查看分析图 go tool pprof -png profile文件  (-png 图片 -text 文档
	// 查看trace go tool trace trace文件
	// see: https://github.com/gin-contrib/pprof
	if app.Conf.EnablePprof {
		pprof.Register(g)
	}

	// HealthCheck 健康检查路由
	g.GET("/health", app.HealthCheck) // 通过 grafana 可视化查看 prometheus 的监控数据，使用插件6671查看
	g.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return g
}

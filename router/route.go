/**
 * @author jiangshangfang
 * @date 2021/8/8 4:53 PM
 **/
package router

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	var Router = gin.Default()

	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		setApiRouter(PublicGroup)
	}

	//PrivateGroup := Router.Group(""){
	//
	//}
	return Router
}

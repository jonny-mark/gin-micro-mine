/**
 * @author jiangshangfang
 * @date 2021/8/8 5:20 PM
 **/
package router

import (
	"github.com/gin-gonic/gin"
	"gin/api/controller"
)

func setApiRouter(router * gin.RouterGroup)  {
	ApiRouter := router.Group("api")
	{
		ApiRouter.GET("test",controller.CreateApi)
	}
}
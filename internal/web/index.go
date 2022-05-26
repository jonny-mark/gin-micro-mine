package web

import (
	"gin/pkg/app"
	"github.com/gin-gonic/gin"
)

// Index home page
func Index(c *gin.Context) {
	var res app.Response

	res.Success(c, "web连接测试")
}

func SignText(c *gin.Context)  {
	var res app.Response

	res.Success(c, "sign解密成功")
}
package web

import (
	"github.com/gin-gonic/gin"
	"github.com/jonny-mark/gin-micro-mine/pkg/app"
	"github.com/jonny-mark/gin-micro-mine/pkg/log"
)

// Index home page
func Index(c *gin.Context) {
	var res app.Response

	res.Success(c, "web连接测试")
}

func SignText(c *gin.Context) {
	//var params map[string]interface{}
	var res app.Response

	err := c.Request.ParseForm()
	if err != nil {
		log.Errorf("etcd unlock the lock err, err: %s", err.Error())
		res.Error(c, err)
	}

	res.Success(c, c.Request.PostForm)
}

/**
 * @author jiangshangfang
 * @date 2021/8/8 5:28 PM
 **/
package controller

import (
	"github.com/gin-gonic/gin"
)

func CreateApi(c *gin.Context) {
	c.Writer.Write([]byte("this is page 2"))
}


package web

import (
	"gin/pkg/app"
	"github.com/gin-gonic/gin"
)

// Index home page
func Index(c *gin.Context) {
	var res app.Response

	res.Success(c, "ffffff")
}

/**
 * @author jiangshangfang
 * @date 2021/12/12 5:30 PM
 **/
package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//不设置缓存
func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

//OPTIONS设置
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*") //允许访问所有域
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}

//安全性
func Secure(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	//不允许在frame中展示
	c.Header("X-Frame-Options", "DENY")
	c.Header("X-Content-Type-Options", "nosniff")
	//当检测到反射的XSS攻击时阻止加载页面
	c.Header("X-XSS-Protection", "1; mode=block")
	//只能通过HTTPS访问
	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}
}

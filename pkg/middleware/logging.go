/**
 * @author jiangshangfang
 * @date 2021/12/12 5:40 PM
 **/
package middleware

import (
	"bytes"
	"encoding/json"
	"gin-micro-mine/pkg/app"
	"gin-micro-mine/pkg/errcode"
	"gin-micro-mine/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/willf/pad"
	"regexp"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLogWriter) Write(data []byte) (n int, err error) {
	w.body.Write(data)
	return w.ResponseWriter.Write(data)
}

//对每一次请求记录日志
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()

		path := c.Request.URL.Path
		//登陆不记录
		reg := regexp.MustCompile("(/login)")
		if reg.MatchString(path) {
			return
		}

		method := c.Request.Method
		ip := c.ClientIP()
		blw := &bodyLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = blw

		c.Next()

		//记录基础信息
		end := time.Now().UTC()
		latency := end.Sub(start)

		var code int
		var message string
		var response app.Response
		if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
			log.Errorf("response body can not unmarshal to config.Response struct, body: `%s`, err: %v",
				blw.body.Bytes(), err)
			code = errcode.ServerError.Code
			message = err.Error()
		} else {
			code = response.Code
			message = response.Message
		}

		//-13s长度最小为13个字符，不足的话右边补空格 pad.Right
		log.Infof("%-13s | %-12s | %s %s | %d | {code: %d, message: %s}", latency, ip, pad.Right(method, 5, ""), path, blw.Status(), code, message)
	}
}

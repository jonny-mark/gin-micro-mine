/**
 * @author jiangshangfang
 * @date 2021/10/27 9:17 PM
 **/
package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin/pkg/utils"
	"gin/pkg/errcode"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success return a success response
func (r *Response) Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	c.JSON(http.StatusOK, Response{
		Code: errcode.Success.GetCode(),
		Msg:  errcode.Success.GetMsg(),
		Data: data,
	})
}

func (r *Response) Error(c *gin.Context, err error) {
	if err == nil {
		c.JSON(http.StatusOK, Response{
			Code: errcode.Success.GetCode(),
			Msg:  errcode.Success.GetMsg(),
			Data: gin.H{},
		})
	}
	if v, ok := err.(*errcode.Error); ok {
		response := Response{
			Code: v.GetCode(),
			Msg:  v.GetMsg(),
			Data: gin.H{},
		}
		c.JSON(v.StatusCode(), response)
		return
	}
}

// SendResponse 返回json
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errcode.DecodeErr(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  message,
		Data: data,
	})
}

// RouteNotFound 未找到相关路由
func RouteNotFound(c *gin.Context) {
	c.String(http.StatusNotFound, "the route not found")
}

// healthCheckResponse 健康检查响应结构体
type healthCheckResponse struct {
	Status   string `json:"status"`
	Hostname string `json:"hostname"`
}

// HealthCheck will return OK if the underlying BoltDB is healthy. At least healthy enough for demoing purposes.
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, healthCheckResponse{Status: "UP", Hostname: utils.GetHostname()})
}

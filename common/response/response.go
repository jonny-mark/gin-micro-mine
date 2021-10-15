package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin/common/constant"
	"go-eagle/pkg/errcode"
	"go-eagle/pkg/utils"
)

var c *gin.Context

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ErrResult(err error, data interface{}) {
	code, msg := errcode.DecodeErr(err)
	c.JSON(http.StatusOK, Response{code, msg, data})
}

func Result(code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{code, msg, data})
}

func Ok() {
	Result(constant.SUCCESS, "操作成功", map[string]interface{}{})
}

func OkWithData(data interface{}) {
	Result(constant.SUCCESS, "操作成功", data)
}

func Fail(code int) {
	Result(code, zhCNText[code], map[string]interface{}{})
}

func FailWithData(code int, data interface{}) {
	Result(code, zhCNText[code], data)
}

// healthCheckResponse 健康检查响应结构体
type healthCheckResponse struct {
	Status   string `json:"status"`
	Hostname string `json:"hostname"`
}

// HealthCheck will return OK if the underlying BoltDB is healthy. At least healthy enough for demoing purposes.
func HealthCheck() {
	c.JSON(http.StatusOK, healthCheckResponse{Status: "UP", Hostname: utils.GetHostname()})
}
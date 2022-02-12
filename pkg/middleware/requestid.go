/**
 * @author jiangshangfang
 * @date 2022/1/24 7:41 PM
 **/
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	HeaderXRequestIDKey = "X-Request-Id"
	ContextRequestIDKey = "x_request_id"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(HeaderXRequestIDKey)
		if requestID == "" {
			requestID = generateID()
			c.Request.Header.Set(HeaderXRequestIDKey, requestID)
		}
		// 在应用上暴露requestID
		c.Set(ContextRequestIDKey, requestID)
		// 设置X-Request-ID响应头
		c.Writer.Header().Set(HeaderXRequestIDKey, requestID)
		c.Next()
	}
}

//随机字符串，eg: 76d27e8c-a80e-48c8-ad20-e5562e0f67e4
func generateID() string {
	reqID, _ := uuid.NewRandom()
	return reqID.String()
}

/**
 * @author jiangshangfang
 * @date 2021/12/12 5:40 PM
 **/
package middleware

import "github.com/gin-gonic/gin"

//对每一次请求记录日志
func Logging() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
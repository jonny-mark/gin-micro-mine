package middleware

import (
	"github.com/gin-gonic/gin"
	"gin-micro-mine/pkg/app"
	"gin-micro-mine/pkg/errcode"
)

// Auth authorize user
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		ctx, err := app.ParseRequest(c)
		if err != nil {
			app.NewResponse().Error(c, errcode.AuthorizationError)
			c.Abort()
			return
		}

		// set uid to context
		c.Set("uid", ctx.UserID)

		c.Next()
	}
}
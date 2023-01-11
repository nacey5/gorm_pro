package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"time"
)

// 时间控制中间件
func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

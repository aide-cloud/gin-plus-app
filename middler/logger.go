package middler

import (
	"log"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware 日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		c.Next()
		// after request
		log.Printf("[LoggerMiddleware] path: %s, status: %d", path, c.Writer.Status())
	}
}

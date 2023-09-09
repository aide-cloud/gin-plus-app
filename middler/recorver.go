package middler

import (
	"log"

	"github.com/gin-gonic/gin"
)

// RecoverMiddleware recover中间件
func RecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatus(500)
			}
		}()
		c.Next()
		log.Println("[RecoverMiddleware] panic recover")
	}
}

package middler

import (
	"log"

	"github.com/gin-gonic/gin"
)

// MustLogin 必须登录
func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("[MustLogin] must login")
	}
}

package article

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (a *Article) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("create article")
	}
}

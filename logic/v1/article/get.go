package article

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (a *Article) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("get article")
	}
}

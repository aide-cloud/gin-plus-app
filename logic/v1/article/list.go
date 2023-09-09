package article

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (a *Article) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("list article")
	}
}

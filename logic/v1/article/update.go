package article

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (a *Article) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("update article")
	}
}

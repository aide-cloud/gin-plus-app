package v1

import (
	"log"

	"gin-plus-app/logic/v1/article"
	"gin-plus-app/logic/v1/tag"
	"github.com/gin-gonic/gin"
)

type V1 struct {
	Article *article.Article
	Tag     *tag.Tag
}

func New() *V1 {
	return &V1{
		Article: article.NewArticleAPI(),
		Tag:     tag.NewTagAPI(),
	}
}

func (v *V1) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			log.Println("[v1] middleware 1")
		},
	}
}

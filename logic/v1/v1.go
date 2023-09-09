package v1

import (
	"gin-plus-app/logic/v1/article"
	"github.com/gin-gonic/gin"
)

type V1 struct {
	Article *article.Article
}

func New() *V1 {
	return &V1{
		Article: article.NewArticleAPI(),
	}
}

func (v *V1) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {

		},
	}
}

func (v *V1) MethoderMiddlewares() map[string][]gin.HandlerFunc {
	return map[string][]gin.HandlerFunc{
		"Update": {
			func(ctx *gin.Context) {

			},
			func(ctx *gin.Context) {

			},
		},
	}
}

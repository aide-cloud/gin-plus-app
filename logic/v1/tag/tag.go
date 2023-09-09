package tag

import (
	"log"

	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

var _ ginplus.Middlewarer = (*Tag)(nil)

type Tag struct {
}

func NewTagAPI() *Tag {
	return &Tag{}
}

func (t *Tag) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			ctx.Set("tag", "tag middleware")
			log.Println("[tag] middleware 1")
		},
	}
}

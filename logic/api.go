package logic

import (
	"log"

	"gin-plus-app/logic/enterprise"
	"gin-plus-app/logic/user"
	v1 "gin-plus-app/logic/v1"
	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

var _ ginplus.Middlewarer = (*Api)(nil)

type Api struct {
	V1         *v1.V1
	User       *user.User
	Enterprise *enterprise.Enterprise
}

func (a *Api) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			log.Println("[Api] middleware 1")
		},
		func(ctx *gin.Context) {
			log.Println("[Api] middleware 2")
		},
	}
}

func NewApi() *Api {
	return &Api{
		V1:         v1.New(),
		User:       user.NewUserAPI(),
		Enterprise: enterprise.New(),
	}
}

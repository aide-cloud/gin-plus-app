package user

import (
	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

type User struct {
	// 放一些log或者下级路由model
}

func (u *User) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {

		},
	}
}

func (u *User) MethoderMiddlewares() map[string][]gin.HandlerFunc {
	return map[string][]gin.HandlerFunc{
		"Update": {
			func(ctx *gin.Context) {

			},
			func(ctx *gin.Context) {

			},
		},
	}
}

func NewUserAPI() *User {
	return &User{}
}

var _ ginplus.MethoderMiddlewarer = (*User)(nil)
var _ ginplus.Middlewarer = (*User)(nil)

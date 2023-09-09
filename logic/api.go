package logic

import (
	"gin-plus-app/logic/user"
	v1 "gin-plus-app/logic/v1"
)

type Api struct {
	V1   *v1.V1
	User *user.User
}

func NewApi() *Api {
	return &Api{
		V1:   v1.New(),
		User: user.NewUserAPI(),
	}
}

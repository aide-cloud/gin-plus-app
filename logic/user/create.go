package user

import (
	"context"

	"gin-plus-app/model"
)

type (
	CreateReq struct {
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Avatar string `json:"avatar"`
		Remark string `json:"remark"`
	}

	CreateResp struct {
		ID uint `json:"id"`
	}
)

// Create 创建用户
func (u *User) Create(ctx context.Context, req *CreateReq) (*CreateResp, error) {
	userModel := model.User{
		Name:   req.Name,
		Age:    req.Age,
		Avatar: req.Avatar,
		Remark: req.Remark,
	}

	//if err := model.NewUserDB(ctx).Create(&userModel).Error; err != nil {
	//	return nil, err
	//}

	return &CreateResp{
		ID: userModel.ID,
	}, nil
}

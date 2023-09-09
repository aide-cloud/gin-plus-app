package user

import (
	"context"

	"gin-plus-app/model"
)

type (
	UpdateReq struct {
		ID     uint   `uri:"id"`
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Avatar string `json:"avatar"`
		Remark string `json:"remark"`
	}

	UpdateResp struct {
		ID uint `json:"id"`
	}
)

// Update 创建用户
func (u *User) Update(ctx context.Context, req *UpdateReq) (*UpdateResp, error) {
	userModel := model.User{
		Name:   req.Name,
		Age:    req.Age,
		Avatar: req.Avatar,
		Remark: req.Remark,
	}

	//if err := model.NewUserDB(ctx).Update(&userModel).Error; err != nil {
	//	return nil, err
	//}

	return &UpdateResp{
		ID: userModel.ID,
	}, nil
}

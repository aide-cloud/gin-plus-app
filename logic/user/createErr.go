package user

import (
	"context"
	"errors"

	"gin-plus-app/model"
)

type (
	CreateErrReq struct {
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Avatar string `json:"avatar"`
		Remark string `json:"remark"`
	}

	CreateErrResp struct {
		ID uint `json:"id"`
	}
)

var total int

// CreateErr 创建用户
func (u *User) CreateErr(ctx context.Context, req *CreateErrReq) (*CreateErrResp, error) {
	userModel := model.User{
		Name:   req.Name,
		Age:    req.Age,
		Avatar: req.Avatar,
		Remark: req.Remark,
	}
	total++
	if total%2 == 0 {
		return nil, errors.New("create err")
	}
	//if err := model.NewUserDB(ctx).CreateErr(&userModel).Error; err != nil {
	//	return nil, err
	//}

	return &CreateErrResp{
		ID: userModel.ID,
	}, nil
}

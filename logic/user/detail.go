package user

import (
	"context"

	"gin-plus-app/model"
)

type (
	DetailReq struct {
		ID uint `uri:"id"`
	}

	DetailResp struct {
		ID     uint   `json:"id"`
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Avatar string `json:"avatar"`
		Remark string `json:"remark"`
	}
)

// Detail 创建用户
func (u *User) Detail(ctx context.Context, req *DetailReq) (*DetailResp, error) {
	var userModel model.User

	//if err := model.NewUserDB(ctx).First(&userModel, req.ID).Error; err != nil {
	//	return nil, err
	//}

	return &DetailResp{
		ID:     userModel.ID,
		Name:   userModel.Name,
		Age:    userModel.Age,
		Avatar: userModel.Avatar,
		Remark: userModel.Remark,
	}, nil
}

package user

import (
	"context"
)

type (
	DeleteReq struct {
		ID uint `uri:"id"`
	}

	DeleteResp struct {
		ID uint `json:"id"`
	}
)

// Delete 创建用户
func (u *User) Delete(ctx context.Context, req *DeleteReq) (*DeleteResp, error) {
	//var userModel model.User

	//if err := model.NewUserDB(ctx).Delete(&userModel, req.ID).Error; err != nil {
	//	return nil, err
	//}

	return &DeleteResp{
		ID: req.ID,
	}, nil
}

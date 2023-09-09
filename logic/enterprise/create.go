package enterprise

import (
	"context"
)

type (
	CreateReq struct {
		Name string `json:"name" binding:"required"`
	}
	CreateResp struct {
		Id string `json:"id"`
	}
)

func (e *Enterprise) Create(ctx context.Context, req *CreateReq) (*CreateResp, error) {
	return &CreateResp{
		Id: "1",
	}, nil
}

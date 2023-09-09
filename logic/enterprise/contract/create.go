package contract

import (
	"context"
)

type CreateReq struct {
	EID  string `uri:"eid" skip:"true"`
	Name string `json:"name" binding:"required"`
}

type CreateResp struct {
	Id  string `json:"id"`
	EID string `json:"eid"`
}

func (c *Contract) PostInfo(ctx context.Context, req *CreateReq) (*CreateResp, error) {
	return &CreateResp{
		Id:  "1",
		EID: req.EID,
	}, nil
}

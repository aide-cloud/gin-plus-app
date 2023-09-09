package contract

import (
	"context"
)

type (
	// UpdateReq 定义请求参数
	UpdateReq struct {
		EID  string `uri:"eid" skip:"true"`
		ID   string `uri:"id"`
		Name string `json:"name"`
	}

	// UpdateResp 定义返回参数
	UpdateResp struct {
		Id  string `json:"id"`
		EID string `json:"eid"`
	}
)

// UpdateInfo 定义更新方法
func (c *Contract) UpdateInfo(ctx context.Context, req *UpdateReq) (*UpdateResp, error) {
	return &UpdateResp{
		Id:  req.ID,
		EID: req.EID,
	}, nil
}

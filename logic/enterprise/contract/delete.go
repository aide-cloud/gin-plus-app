package contract

import (
	"context"
)

type (
	// DeleteReq 定义请求参数
	DeleteReq struct {
		EID string `uri:"eid" skip:"true"`
		ID  string `uri:"id"`
	}

	// DeleteResp 定义返回参数
	DeleteResp struct {
		Id  string `json:"id"`
		EID string `json:"eid"`
	}
)

// DeleteInfo 定义更新方法
func (c *Contract) DeleteInfo(ctx context.Context, req *DeleteReq) (*DeleteResp, error) {
	return &DeleteResp{
		Id:  req.ID,
		EID: req.EID,
	}, nil
}

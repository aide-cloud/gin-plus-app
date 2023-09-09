package contract

import (
	"context"
)

type (
	// DetailReq 定义请求参数
	DetailReq struct {
		EID string `uri:"eid" skip:"true"`
		ID  string `uri:"id"`
	}

	// DetailResp 定义返回参数
	DetailResp struct {
		Id   string `json:"id"`
		EID  string `json:"eid"`
		Name string `json:"name"`
	}
)

// DetailInfo 定义更新方法
func (c *Contract) DetailInfo(ctx context.Context, req *DetailReq) (*DetailResp, error) {
	return &DetailResp{
		Id:   req.ID,
		Name: "test",
		EID:  req.EID,
	}, nil
}

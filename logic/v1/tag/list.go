package tag

import (
	"context"
)

type (
	// ListTagRequest struct
	ListTagRequest struct {
		Size int `form:"size"`
		Page int `form:"page"`
	}
	// ListTagResponse struct
	ListTagResponse struct {
		List  []*Item `json:"list"`
		Total int64   `json:"total"`
	}

	// Item struct
	Item struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	}
)

// List 创建标签
func (t *Tag) List(ctx context.Context, req *ListTagRequest) (*ListTagResponse, error) {
	return &ListTagResponse{
		List: []*Item{
			{
				ID:    1,
				Name:  "test",
				Color: "red",
			},
		},
		Total: 1,
	}, nil
}

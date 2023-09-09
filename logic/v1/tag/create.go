package tag

import (
	"context"
)

type (
	// CreateTagRequest struct
	CreateTagRequest struct {
		Name  string `json:"name" binding:"required"`
		Color string `json:"color" binding:"required"`
	}
	// CreateTagResponse struct
	CreateTagResponse struct {
		ID uint `json:"id"`
	}
)

// Create 创建标签
func (t *Tag) Create(ctx context.Context, req *CreateTagRequest) (*CreateTagResponse, error) {
	return &CreateTagResponse{
		ID: 1,
	}, nil
}

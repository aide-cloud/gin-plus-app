package user

import (
	"context"
	"time"

	"gin-plus-app/model"
)

type (
	ListReq struct {
		Current int    `form:"current"`
		Size    int    `form:"size"`
		Keyword string `form:"keyword"`
	}

	ListResp struct {
		List    []Item `json:"list"`
		Total   int64  `json:"total"`
		Current int    `json:"current"`
		Size    int    `json:"size"`
	}

	Item struct {
		ID     uint   `json:"id"`
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Avatar string `json:"avatar"`
		Remark string `json:"remark"`
	}
)

// List 创建用户
func (u *User) List(ctx context.Context, req *ListReq) (*ListResp, error) {
	var userModelList []model.User
	var total int64

	time.Sleep(1 * time.Second)

	//db := model.NewUserDB(ctx)
	//
	//if req.Keyword != "" {
	//	db = db.Where("name like ?", "%"+req.Keyword+"%")
	//}

	//db.Count(&total)
	//
	//db = db.Offset((req.Current - 1) * req.Size).Limit(req.Size).Find(&userModelList)
	//
	//if err := db.Find(&userModelList).Error; err != nil {
	//	return nil, err
	//}

	list := make([]Item, 0, len(userModelList))
	for _, userModel := range userModelList {
		list = append(list, Item{
			ID:     userModel.ID,
			Name:   userModel.Name,
			Age:    userModel.Age,
			Avatar: userModel.Avatar,
			Remark: userModel.Remark,
		})
	}

	return &ListResp{
		List:    list,
		Total:   total,
		Current: req.Current,
		Size:    req.Size,
	}, nil
}

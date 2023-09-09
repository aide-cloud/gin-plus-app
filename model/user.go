package model

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Name   string `gorm:"column:name;NOT NULL;comment:姓名;type:varchar(32)" json:"name"`
	Age    int    `gorm:"column:age;NOT NULL;comment:年龄;type:int(11)" json:"age"`
	Avatar string `gorm:"column:avatar;NOT NULL;comment:头像;type:varchar(255)" json:"avatar"`
	Remark string `gorm:"column:remark;NOT NULL;comment:备注;type:varchar(255)" json:"remark"`
}

func NewUserDB(ctx context.Context, txs ...*gorm.DB) *gorm.DB {
	return GetDB(txs...).WithContext(ctx).Model(User{})
}

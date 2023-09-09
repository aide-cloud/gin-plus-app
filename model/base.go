package model

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
}

type EModel struct {
	BaseModel
	EID string `gorm:"<-:create;column:eid;NOT NULL;index;comment:企业id" json:"eid"`
}

var db *gorm.DB

func GetDB(dbs ...*gorm.DB) *gorm.DB {
	if len(dbs) > 0 {
		return db
	}
	return db
}

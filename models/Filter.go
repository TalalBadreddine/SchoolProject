package models

import (
	"github.com/jinzhu/gorm"
)

type Filter struct {
	Class   string `json:"class"`
	Student string `json:"student"`
	Teacher string `json:"teacher"`
}

func FilterByClass(class *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if *class == "" {
			return db
		}
		//db.Joins("classes").Where("class_id = ?", class)
		return db.Where("classs =?", class)
	}
}

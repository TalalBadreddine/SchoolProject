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
		return db.Where("classs =?", class)
	}
}

func FilterByClasses(class []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nil == class || len(class) == 0 || class[0] == "" {
			return db
		}
		return db.Where("classs IN (?)", class)
	}
}

func FilterByTeachers(teachers []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nil == teachers || len(teachers) == 0 || teachers[0] == "" {
			return db
		}
		return db.Where("teacher IN (?)", teachers)
	}
}

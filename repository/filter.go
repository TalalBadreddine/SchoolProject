package repository

import (
	"gorm.io/gorm"
)

type Filter struct {
	Class    string `json:"class"`
	Student  string `json:"student"`
	Teacher  string `json:"teacher"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}

func FilterByClass(class *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if *class == "" {
			return db
		}
		return db.Where("classs = ?", class)
	}
}

func FilterByClasses(class []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nil == class || len(class) == 0 || class[0] == "" {
			return db
		}
		return db.Where("FirstName IN (?)", class)
	}
}

func FilterByStudents(student []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nil == student || len(student) == 0 || student[0] == "" {
			return db
		}
		return db.Where("classs IN (?)", student)
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

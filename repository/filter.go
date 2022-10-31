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

//TODO move all scopes to entity
func FilterByClassesId(class []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nil == class || len(class) == 0 || class[0] == "" {
			return db
		}
		return db.Where("classes.id IN (?)", class)
	}
}

func FilterByStudentsId(student []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nil == student || len(student) == 0 || student[0] == "" {
			return db
		}
		return db.Where("students.id IN (?)", student)
	}
}

func FilterByTeachers(teachers []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nil == teachers || len(teachers) == 0 || teachers[0] == "" {
			return db
		}
		return db.Where("teachers.id IN (?)", teachers)
	}
}

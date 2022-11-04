package entity

import (
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Subject         string             `json:"subject"`
	Code            string             `json:"code"`
	Teachers        []*Teacher         `json:"teachers" gorm:"many2many:teacher_classes; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Students        []*Student         `json:"students" gorm:"many2many:students_classes; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StudentsClasses []*StudentsClasses `json:"grade" gorm:"one2many:students_classes;"`
	HighestGrade    int                `gorm:"-:migration;column:highest_grade"`
}

func FilterByClassesId(class []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nil == class || len(class) == 0 || class[0] == "" {
			return db
		}
		return db.Where("classes.id IN (?)", class)
	}
}

func FilterByClassesCode(codes []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nil == codes || len(codes) == 0 || codes[0] == "" {
			return db
		}
		return db.Where("classes.code IN (?)", codes)
	}
}

func FilterByClassesSubject(subjects []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nil == subjects || len(subjects) == 0 || subjects[0] == "" {
			return db
		}
		return db.Where("classes.subject IN (?)", subjects)
	}
}

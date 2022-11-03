package entity

import (
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Subject  string             `json:"subject"`
	Code     string             `json:"code"`
	Teachers []*Teacher         `json:"teachers" gorm:"many2many:teacher_classes; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Students []*Student         `json:"students" gorm:"many2many:students_classes; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Grade    []*StudentsClasses `json:"grade" gorm:"one2many:students_classes;foreignKey:Grade;"`
}

func FilterByClassesId(class []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nil == class || len(class) == 0 || class[0] == "" {
			return db
		}
		return db.Where("classes.id IN (?)", class)
	}
}

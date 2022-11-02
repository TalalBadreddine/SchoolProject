package entity

import (
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Subject  string
	Code     string
	Teachers []*Teacher `gorm:"many2many:teacher_classes; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Students []*Student `gorm:"many2many:students_classes; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func FilterByClassesId(class []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nil == class || len(class) == 0 || class[0] == "" {
			return db
		}
		return db.Where("classes.id IN (?)", class)
	}
}

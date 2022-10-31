package entity

import (
	"github.com/jinzhu/gorm"
)

type Class struct {
	gorm.Model
	Subject  string
	Code     string
	Teachers []*Teacher `gorm:"many2many:teacher_classes; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Students []*Student `gorm:"many2many:students_classes; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

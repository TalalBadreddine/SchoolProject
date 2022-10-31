package entity

import (
	"github.com/jinzhu/gorm"
)

type Teacher struct {
	gorm.Model
	FirstName string   `json:"firstName" gorm:"column:first_name"`
	LastName  string   `json:"lastName" gorm:"column:last_name"`
	Classes   []*Class `gorm:"many2many:teacher_classes; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type TeacherClassRequest struct {
	TeacherId int `json:"teacherId"`
	ClassId   int `json:"classId"`
}

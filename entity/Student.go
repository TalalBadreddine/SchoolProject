package entity

import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	FirstName string   `json:"firstName" gorm:"column:first_name"`
	LastName  string   `json:"lastName" gorm:"column:last_name"`
	Classes   []*Class `json:"studentsClasses" gorm:"many2many:students_classes; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type StudentClassRequest struct {
	StudentId int `json:"studentId"`
	ClassId   int `json:"classId"`
}

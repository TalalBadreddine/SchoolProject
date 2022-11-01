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

func FilterByTeachers(teachers []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nil == teachers || len(teachers) == 0 || teachers[0] == "" {
			return db
		}
		return db.Where("teachers.id IN (?)", teachers)
	}
}

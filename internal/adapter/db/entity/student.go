package entity

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	FirstName      string             `json:"firstName" gorm:"column:first_name"`
	LastName       string             `json:"lastName" gorm:"column:last_name"`
	Classes        []*Class           `json:"studentsClasses" gorm:"many2many:students_classes; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StudentClasses []*StudentsClasses `json:"studentClasses" gorm:"one2many:students_classes; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func FilterByStudentsId(student []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nil == student || len(student) == 0 || student[0] == "" {
			return db
		}
		return db.Where("students.id IN (?)", student)
	}
}

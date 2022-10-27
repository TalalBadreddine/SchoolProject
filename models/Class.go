package models

import (
	"fmt"
	"server/storage"

	"github.com/jinzhu/gorm"
)

type Class struct {
	gorm.Model
	Subject  string
	Code     string
	Teachers []*Teacher `gorm:"many2many:teacherClasses; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Students []*Student `gorm:"many2many:studentsClasses; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func AddClass(class Class) string {
	var db = storage.GetDBInstance()

	err := db.Create(&class)
	message := fmt.Sprintf("Class With Subject: %v %v is added", class.Subject, class.Code)

	if err.Error != nil {
		message = fmt.Sprintf("Error: %v", err.Error.Error())
	}

	return message
}

// TODO: if class doest not exist it's getting greated
func GetClassById(classId int) Class {
	var db = storage.GetDBInstance()

	var class Class

	db.First(&class, classId)

	return class
}

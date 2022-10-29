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
	message := "Added " + DisplayClass(class)

	if err.Error != nil {
		message = fmt.Sprintf("Error: %v", err.Error.Error())
	}

	return message
}

func DisplayClass(class Class) string {
	return fmt.Sprintf("ID: %v ,Subject: %v , Class Code: %v \n", class.ID, class.Subject, class.Code)
}

func GetClassById(classId int) Class {
	var db = storage.GetDBInstance()

	var class Class

	db.First(&class, classId)

	return class
}

func GetAllClasses() []*Class {
	var db = storage.GetDBInstance()
	var classes []*Class

	db.Find(&classes)

	return classes
}

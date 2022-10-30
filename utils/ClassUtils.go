package utils

import (
	"fmt"
	"server/entity"
	"server/storage"
)

func AddClass(class entity.Class) *entity.Class {
	var db = storage.GetDBInstance()

	err := db.Create(&class)

	if err.Error != nil {
		return nil
	}

	return &class
}

func GetClassById(classId int) entity.Class {
	var db = storage.GetDBInstance()

	var class entity.Class

	db.First(&class, classId)

	return class
}

func GetAllClasses() []*entity.Class {
	var db = storage.GetDBInstance()
	var classes []*entity.Class

	db.Find(&classes)

	return classes
}

func DisplayClass(class entity.Class) string {
	return fmt.Sprintf("ID: %v ,Subject: %v , Class Code: %v \n", class.ID, class.Subject, class.Code)
}

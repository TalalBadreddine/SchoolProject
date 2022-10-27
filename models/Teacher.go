package models

import (
	"fmt"
	"server/storage"

	"github.com/jinzhu/gorm"
)

type Teacher struct {
	gorm.Model
	FirstName string `json:"firstName" gorm:"column:first_name"`
	LastName  string `json:"lastName" gorm:"column:last_name"`
	Classes   []*Class
}

type TeacherClassRequest struct {
	TeacherId int `json:"teacherId"`
	ClassId   int `json:"classId"`
}

func AddTeacher(teacher Teacher) string {
	var db = storage.GetDBInstance()

	err := db.Create(&teacher)
	var response = fmt.Sprintf("Added Teacher -> ID: %v, FirstName: %v, LastName: %v", teacher.ID, teacher.FirstName, teacher.LastName)

	if err.Error != nil {
		response = fmt.Sprintf("Error while adding Teacher: %v", err.Error.Error())
	}

	return response
}

func AddTeacherInClass(teacherId int, classId int) string {
	var db = storage.GetDBInstance()
	var class = GetClassById(classId)
	var teacher Teacher

	db.First(&teacher, teacherId)

	teacher.Classes = append(teacher.Classes, &class)
	db.Save(&teacher)

	return "Updated"
}

func GetTeacherById(teacherId int) Teacher {
	var db = storage.GetDBInstance()

	var teacher Teacher

	db.First(&teacher, teacherId)

	return teacher
}

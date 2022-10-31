package repository

import (
	"fmt"
	"server/entity"
	"server/filter"
	"server/storage"
	"strings"
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

func SearchClasses(filter filter.ClassFilter) []*entity.Class {
	var db = storage.GetDBInstance()
	var classes []*entity.Class

	if filter.PerPage == 0 {
		filter.PerPage = 5
	}

	offset := (filter.Page) * filter.PerPage

	var studentArray = strings.Split(filter.StudentId, ",")
	//var teacherArray = strings.Split(filter.Teacher, ",")

	db.Preload("Teachers").
		Preload("Students").
		Joins("LEFT JOIN students_classes on students_classes.class_id = classes.id").
		Joins("JOIN students on students.id = students_classes.student_id").
		Joins("LEFT JOIN teacher_classes on teacher_classes.class_id = classes.id").
		Joins("JOIN teachers on teachers.id = teacher_classes.teacher_id").
		Scopes(FilterByStudentsId(studentArray)).
		Offset(offset).
		Limit(filter.PerPage).
		Find(&classes)

	return classes
}

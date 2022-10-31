package repository

import (
	"fmt"
	"server/entity"
	"server/storage"
	"strings"
)

func AddTeacher(teacher entity.Teacher) *entity.Teacher {
	var db = storage.GetDBInstance()

	err := db.Create(&teacher)

	if err.Error != nil {
		return nil
	}

	return &teacher
}

func DisplayTeacher(teacher entity.Teacher) string {
	return fmt.Sprintf("Teacher -> ID: %v, FirstName: %v, LastName: %v \n", teacher.ID, teacher.FirstName, teacher.LastName)
}

func AddTeacherInClass(teacherId int, classId int) *entity.Teacher {
	var db = storage.GetDBInstance()
	var class = GetClassById(classId)
	var teacher entity.Teacher

	db.First(&teacher, teacherId)

	teacher.Classes = append(teacher.Classes, &class)
	db.Save(&teacher)

	return &teacher
}

func GetTeacherById(teacherId int) entity.Teacher {
	var db = storage.GetDBInstance()

	var teacher entity.Teacher

	db.First(&teacher, teacherId)

	return teacher
}

func GetAllTeachers() []*entity.Teacher {
	var db = storage.GetDBInstance()

	var teacher []*entity.Teacher

	db.Find(&teacher)

	return teacher
}

func SearchTeachers(filter Filter) []*entity.Teacher {
	var db = storage.GetDBInstance()
	var teachers []*entity.Teacher

	if filter.PageSize == 0 {
		filter.PageSize = 5
	}

	offset := (filter.Page - 1) * filter.PageSize

	var classArray = strings.Split(filter.Class, ",")
	var studentArray = strings.Split(filter.Student, ",")

	db.Preload("Classes").
		Preload("Classes.Teachers").
		Preload("Classes.Students").
		Scopes(FilterByClassesId(classArray), FilterByStudentsId(studentArray)).
		Offset(offset).
		Limit(filter.PageSize).
		Find(&teachers)

	return teachers
}
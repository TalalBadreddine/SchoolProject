package utils

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

func SearchTeachers(filter entity.Filter) []*entity.Teacher {
	var db = storage.GetDBInstance()
	var teachers []*entity.Teacher

	offset := (filter.Page - 1) * filter.PageSize

	var classArray = strings.Split(filter.Class, ",")
	var studentArray = strings.Split(filter.Student, ",")

	db.Limit(filter.PageSize).Offset(offset).Scopes(entity.FilterByClasses(classArray), entity.FilterByStudents(studentArray)).Preload("Classes").Preload("Teachers").Find(&teachers)
	return teachers
}

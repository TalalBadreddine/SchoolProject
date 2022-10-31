package utils

import (
	"log"
	"server/entity"
	"server/storage"
	"strings"
)

func StudentExist(target entity.Student) entity.Student {
	var db = storage.GetDBInstance()

	var student entity.Student
	db.First(&student, target.ID)

	return student
}

func AddStudent(student entity.Student) *entity.Student {
	var db = storage.GetDBInstance()

	err := db.Create(&student)

	if err.Error != nil {
		log.Panic(err.Error.Error())
	}

	return &student
}

func GetAllStudents() []*entity.Student {
	var db = storage.GetDBInstance()

	var students []*entity.Student

	db.Find(&students)

	return students
}

func AddStudentInClass(studentId int, classId int) *entity.Student {
	var db = storage.GetDBInstance()
	var class = GetClassById(classId)
	var student entity.Student

	db.First(&student, studentId)

	student.Classes = append(student.Classes, &class)
	db.Save(&student)

	return &student
}

func GetStudentById(studentId int) entity.Student {
	var db = storage.GetDBInstance()

	var student entity.Student

	db.First(&student, studentId)

	return student
}

func SearchStudents(filter entity.Filter) []*entity.Student {
	var db = storage.GetDBInstance()
	var students []*entity.Student

	offset := (filter.Page * filter.PageSize)
	filter.PageSize = 3

	var classArray = strings.Split(filter.Class, ",")
	var teacherArray = strings.Split(filter.Teacher, ",")

	db.Preload("Classes").
		Preload("Classes.Teachers").
		Scopes(entity.FilterByClasses(classArray), entity.FilterByTeachers(teacherArray)).
		Offset(offset).
		Limit(filter.PageSize).
		Find(&students)

	return students
}

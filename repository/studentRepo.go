package repository

import (
	"log"
	"server/entity"
	"server/filter"
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

func SearchStudents(filter filter.StudentFilter) []*entity.Student {
	var db = storage.GetDBInstance()
	var students []*entity.Student

	if filter.PerPage == 0 {
		filter.PerPage = 5
	}

	offset := (filter.Page - 1) * filter.PerPage

	var classArray = strings.Split(filter.ClassesId, ",")
	var studentArray = strings.Split(filter.StudentsId, ",")
	//var teacherArray = strings.Split(filter.Teacher, ",")

	db.Preload("Classes").
		Preload("Classes.Teachers").
		Preload("Classes.Students").
		Joins("LEFT JOIN students_classes on students_classes.student_id = students.id").
		Joins("JOIN classes on classes.id = students_classes.class_id").
		Scopes(FilterByClassesId(classArray), FilterByStudentsId(studentArray)).
		Offset(offset).
		Limit(filter.PerPage).
		Find(&students)

	return students
}

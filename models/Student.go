package models

import (
	"fmt"
	"server/storage"

	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	FirstName string   `json:"firstName" gorm:"column:first_name"`
	LastName  string   `json:"lastName" gorm:"column:last_name"`
	Classes   []*Class `gorm:"many2many:studentsClasses; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type StudentClassRequest struct {
	StudentId int `json:"studentId"`
	ClassId   int `json:"classId"`
}

func StudentExist(target Student) Student {
	var db = storage.GetDBInstance()

	var student Student
	db.First(&student, target.ID)

	return student
}

func AddStudent(student Student) string {
	var db = storage.GetDBInstance()

	err := db.Create(&student)

	if err.Error != nil {
		return err.Error.Error()
	}

	message := fmt.Sprintf("Student %v %v was succefuly created", student.FirstName, student.LastName)

	return message
}

// TODO: Check if updated before sending response
func AddStudentInClass(studentId int, classId int) string {
	var db = storage.GetDBInstance()
	var class = GetClassById(classId)
	var student Student

	db.First(&student, studentId)

	student.Classes = append(student.Classes, &class)
	db.Save(&student)

	return "Updated"
}

func GetAllStudents() []*Student {
	var db = storage.GetDBInstance()

	var students []*Student

	db.Find(&students)

	return students
}

func SearchStudents(filter Filter) []*Student {
	var db = storage.GetDBInstance()
	var students []*Student

	// var classArray = strings.Split(filter.Class, ",")
	// var teacherArray = strings.Split(filter.Teacher, ",")
	// filterByTeacher(teacherArray)

	db.Scopes(FilterByClass(&filter.Class)).Find(&students)
	return students
}

func GetStudentById(studentId int) Student {
	var db = storage.GetDBInstance()

	var student Student

	db.First(&student, studentId)

	return student
}

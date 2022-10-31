package dto

import (
	"server/entity"
)

type Class struct {
	id       uint
	Subject  string
	Code     string
	Teachers []*Teacher
	Students []*Student
}

type GeneralClass struct {
	id      uint
	Subject string
	Code    string
}

func MapClassDto(class *entity.Class) *Class {
	var teachers []*Teacher
	var students []*Student

	for _, teacher := range class.Teachers {
		teachers = append(teachers, MapTeacherDto(teacher))
	}

	for _, student := range class.Students {
		students = append(students, MapStudentDto(student))
	}

	return &Class{
		id:       class.ID,
		Subject:  class.Subject,
		Code:     class.Code,
		Teachers: teachers,
		Students: students,
	}
}

func MapGeneralClassDto(class *entity.Class) *GeneralClass {
	return &GeneralClass{
		id:      class.ID,
		Subject: class.Subject,
		Code:    class.Code,
	}
}

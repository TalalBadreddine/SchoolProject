package dto

import "server/internal/domain/model"

type Class struct {
	id       uint
	Subject  string
	Code     string
	Teachers []*GeneralTeacher
	Students []*GeneralStudent
}

type GeneralClass struct {
	id      uint
	Subject string
	Code    string
}

func MapClassDto(class *model.Class) *Class {
	var teachers []*GeneralTeacher
	var students []*GeneralStudent

	for _, teacher := range class.Teachers {
		teachers = append(teachers, MapGeneralTeacherDto(teacher))
	}

	for _, student := range class.Students {
		students = append(students, MapGeneralStudentDto(student))
	}

	return &Class{
		id:       class.Id,
		Subject:  class.Subject,
		Code:     class.Code,
		Teachers: teachers,
		Students: students,
	}
}

func MapGeneralClassDto(class *model.Class) *GeneralClass {
	return &GeneralClass{
		id:      class.Id,
		Subject: class.Subject,
		Code:    class.Code,
	}
}

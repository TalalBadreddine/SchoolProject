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
	Grade   int
}

func MapClassDto(class *model.Class) *Class {
	var teachers []*GeneralTeacher
	var students []*GeneralStudent

	for _, teacher := range class.Teachers {
		teachers = append(teachers, MapGeneralTeacherDto(teacher))
	}

	for _, student := range class.Students {
		for _, studentClass := range class.StudentsClasses {
			if student.Id == studentClass.StudentId {
				students = append(students, MapStudentInClass(student, studentClass.Grade))
			}
		}
	}

	return &Class{
		id:       class.Id,
		Subject:  class.Subject,
		Code:     class.Code,
		Teachers: teachers,
		Students: students,
	}
}

func MapClassForStudentDto(class *model.Class, grade int) *GeneralClass {
	return &GeneralClass{
		id:      class.Id,
		Subject: class.Subject,
		Code:    class.Code,
		Grade:   grade,
	}
}

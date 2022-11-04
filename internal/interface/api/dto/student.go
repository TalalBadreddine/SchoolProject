package dto

import (
	"server/internal/domain/model"
)

type Student struct {
	id        uint
	FirstName string
	LastName  string
	Classes   []*GeneralClass
}

type GeneralStudent struct {
	id        uint
	FirstName string
	LastName  string
	Grade     int
}

func MapStudentDto(student *model.Student) *Student {

	var classes []*GeneralClass

	for _, class := range student.Classes {
		for _, studentClass := range student.StudentClasses {
			if class.Id == studentClass.ClassId {
				classes = append(classes, MapClassForStudentDto(class, studentClass.Grade))
			}
		}
	}

	var studentsClasses []*StudentsClasses

	for _, studentClass := range student.StudentClasses {
		studentsClasses = append(studentsClasses, MapToStudentsClassesDto(studentClass))
	}

	return &Student{
		id:        student.Id,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Classes:   classes,
	}
}

func MapStudentInClass(student *model.Student, grade int) *GeneralStudent {
	return &GeneralStudent{
		id:        student.Id,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Grade:     grade,
	}
}

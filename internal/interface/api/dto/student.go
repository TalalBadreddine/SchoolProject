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
}

func MapStudentDto(student *model.Student) *Student {

	var classes []*GeneralClass

	for _, class := range student.Classes {
		classes = append(classes, MapGeneralClassDto(class))
	}

	return &Student{
		id:        student.Id,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Classes:   classes,
	}
}

func MapGeneralStudentDto(student *model.Student) *GeneralStudent {
	return &GeneralStudent{
		id:        student.Id,
		FirstName: student.FirstName,
		LastName:  student.LastName,
	}
}

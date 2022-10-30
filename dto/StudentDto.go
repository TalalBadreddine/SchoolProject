package dto

import "server/entity"

type Student struct {
	id        uint
	FirstName string
	LastName  string
	Classes   []*Class
}

func MapStudentDto(student *entity.Student) *Student {

	var classes []*Class

	for _, class := range student.Classes {
		classes = append(classes, MapClassDto(class))
	}

	return &Student{
		id:        student.ID,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Classes:   classes,
	}
}

package dto

import "server/entity"

type Teacher struct {
	id        uint
	FirstName string
	LastName  string
	Classes   []*Class
}

func MapTeacherDto(teacher *entity.Teacher) *Teacher {
	var classes []*Class

	for _, class := range teacher.Classes {
		classes = append(classes, MapClassDto(class))
	}

	return &Teacher{
		id:        teacher.ID,
		FirstName: teacher.FirstName,
		LastName:  teacher.LastName,
		Classes:   classes,
	}
}

package dto

import "server/internal/domain/model"

type Teacher struct {
	id        uint
	FirstName string
	LastName  string
	Classes   []*Class
}

type GeneralTeacher struct {
	id        uint
	FirstName string
	LastName  string
}

func MapGeneralTeacherDto(teacher *model.Teacher) *GeneralTeacher {
	return &GeneralTeacher{
		id:        teacher.Id,
		FirstName: teacher.FirstName,
		LastName:  teacher.LastName,
	}
}

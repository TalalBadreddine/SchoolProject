package mapper

import (
	"server/internal/adapter/db/entity"
	"server/internal/domain/model"
)

func MapToTeacherModel(s *entity.Teacher) *model.Teacher {
	if nil == s {
		return nil
	}

	return &model.Teacher{
		Id:        s.ID,
		FirstName: s.FirstName,
		LastName:  s.LastName,
		Classes:   MapToClassModelArray(s.Classes),
	}
}

func MapToTeacherModelArray(s []*entity.Teacher) []*model.Teacher {

	if nil == s {
		return nil
	}

	var teachers []*model.Teacher

	for _, teacher := range s {
		teachers = append(teachers, MapToTeacherModel(teacher))
	}

	return teachers
}

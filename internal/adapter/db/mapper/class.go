package mapper

import (
	"server/internal/adapter/db/entity"
	"server/internal/domain/model"
)

func MapToClassModel(s *entity.Class) *model.Class {

	if nil == s {
		return nil
	}

	return &model.Class{
		Id:              s.ID,
		Subject:         s.Subject,
		Code:            s.Code,
		Teachers:        MapToTeacherModelArray(s.Teachers),
		Students:        MapToStudentModelArray(s.Students),
		StudentsClasses: MapToStudentClassArray(s.StudentsClasses),
	}

}

func MapToClassModelArray(s []*entity.Class) []*model.Class {

	if nil == s {
		return nil
	}

	var classes []*model.Class

	for _, class := range s {
		classes = append(classes, MapToClassModel(class))
	}

	return classes
}

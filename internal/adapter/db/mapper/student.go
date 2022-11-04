package mapper

import (
	"server/internal/adapter/db/entity"
	"server/internal/domain/model"
)

func MapToStudentModel(s *entity.Student) *model.Student {
	if nil == s {
		return nil
	}

	return &model.Student{
		Id:             s.ID,
		FirstName:      s.FirstName,
		LastName:       s.LastName,
		Classes:        MapToClassModelArray(s.Classes),
		StudentClasses: MapToStudentClassArray(s.StudentClasses),
	}
}

func MapToStudentModelArray(s []*entity.Student) []*model.Student {
	if nil == s {
		return nil
	}

	var students []*model.Student

	for _, student := range s {
		students = append(students, MapToStudentModel(student))
	}

	return students
}

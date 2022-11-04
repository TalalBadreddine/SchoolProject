package mapper

import (
	"server/internal/adapter/db/entity"
	"server/internal/domain/model"
)

func MapToStudentClass(s *entity.StudentsClasses) *model.StudentsClasses {
	if nil == s {
		return nil
	}

	return &model.StudentsClasses{
		StudentId: s.StudentId,
		ClassId:   s.ClassId,
		Grade:     s.Grade,
	}
}

func MapToStudentClassArray(s []*entity.StudentsClasses) []*model.StudentsClasses {

	if nil == s {
		return nil
	}

	var studentsClasses []*model.StudentsClasses

	for _, studentsClass := range s {
		studentsClasses = append(studentsClasses, MapToStudentClass(studentsClass))
	}

	return studentsClasses
}

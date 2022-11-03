package query

import (
	"server/internal/domain/model"
	"server/internal/domain/model/filter"
	"server/internal/domain/repository"
)

type GetAllStudents struct {
	studentRepository repository.StudentRepository
}

func ProvideGetAllStudents(s repository.StudentRepository) GetAllStudents {
	return GetAllStudents{studentRepository: s}
}

func (s GetAllStudents) Handle(filter filter.StudentFilter) []*model.Student {
	return s.studentRepository.SearchStudents(filter)
}

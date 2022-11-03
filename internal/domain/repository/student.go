package repository

import (
	"server/internal/domain/model"
	"server/internal/domain/model/filter"
)

type StudentRepository interface {
	SearchStudents(filter filter.StudentFilter) []*model.Student
}

package query

import (
	"server/internal/domain/model"
	"server/internal/domain/model/filter"
	"server/internal/domain/repository"
)

type GetAllClasses struct {
	ClassRepository repository.ClassRepository
}

func ProvideGetAllClasses(c repository.ClassRepository) GetAllClasses {
	return GetAllClasses{ClassRepository: c}
}

func (c GetAllClasses) Handle(filter filter.ClassFilter) []*model.Class {
	return c.ClassRepository.SearchClasses(filter)
}

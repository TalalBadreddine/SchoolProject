package repository

import (
	"server/internal/domain/model"
	"server/internal/domain/model/filter"
)

type ClassRepository interface {
	SearchClasses(filter filter.ClassFilter) []*model.Class
}

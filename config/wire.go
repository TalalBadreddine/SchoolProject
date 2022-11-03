//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"server/internal/adapter/db/repository"
	"server/internal/interface/api"
	"server/internal/useCase/query"
)

func WireStudentApi(db *gorm.DB) api.Student {
	wire.Build(
		api.ProvideStudent,
		query.ProvideGetAllStudents,
		repository.ProvideStudentRepository,
	)
	return api.Student{}
}

func WireClassApi(db *gorm.DB) api.Class {
	wire.Build(
		api.ProvideClass,
		query.ProvideGetAllClasses,
		repository.ProvideClassRepository,
	)

	return api.Class{}
}

//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"server/controller"
	"server/repository"
)

func WireStudentApi(db *gorm.DB) controller.Student {
	wire.Build(
		controller.ProvideStudent,
		repository.ProvideStudentRepository,
	)
	return controller.Student{}
}

func WireClassApi(db *gorm.DB) controller.Class {
	wire.Build(
		controller.ProvideClass,
		repository.ProvideClassRepository,
	)

	return controller.Class{}
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package config

import (
	"gorm.io/gorm"
	"server/controller"
	"server/repository"
)

// Injectors from wire.go:

func WireStudentApi(db *gorm.DB) controller.Student {
	studentRepository := repository.ProvideStudentRepository(db)
	student := controller.ProvideStudent(studentRepository)
	return student
}

func WireClassApi(db *gorm.DB) controller.Class {
	classRepository := repository.ProvideClassRepository(db)
	class := controller.ProvideClass(classRepository)
	return class
}

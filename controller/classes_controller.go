package controller

import (
	"net/http"
	"server/dto"
	"server/filter"
	"server/repository"

	"github.com/labstack/echo/v4"
)

type Class struct {
	ClassRepository repository.ClassRepository
}

func ProvideClass(classRepository repository.ClassRepository) Class {
	return Class{ClassRepository: classRepository}
}

func (class Class) GetClasses(c echo.Context) error {

	var classFilter filter.ClassFilter

	err := c.Bind(&classFilter)

	if err != nil {
		return err
	}

	classes := class.ClassRepository.SearchClasses(classFilter)
	var classesDto []*dto.Class

	for _, class := range classes {
		classesDto = append(classesDto, dto.MapClassDto(class))
	}

	return c.JSON(http.StatusOK, classesDto)
}

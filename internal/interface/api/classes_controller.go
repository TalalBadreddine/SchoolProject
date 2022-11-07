package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/internal/domain/model/filter"
	"server/internal/interface/api/dto"
	"server/internal/useCase/query"
)

type Class struct {
	GetAllClasses query.GetAllClasses
}

func ProvideClass(c query.GetAllClasses) Class {
	return Class{GetAllClasses: c}
}

func (class Class) GetClasses(c echo.Context) error {

	var classFilter filter.ClassFilter

	err := c.Bind(&classFilter)

	if err != nil {
		return err
	}

	classes := class.GetAllClasses.Handle(classFilter)
	var classesDto []*dto.Class

	for _, class := range classes {
		classesDto = append(classesDto, dto.MapClassDto(class))
	}

	return c.JSON(http.StatusOK, classesDto)
}

func (class Class) GetStatistics(c echo.Context) error {
	var classFilter filter.ClassFilter

	err := c.Bind(&classFilter)

	if err != nil {
		return err
	}

	classes := class.GetAllClasses.Handle(classFilter)

	var statisticsDto []*dto.Statistics

	for _, class := range classes {
		statisticsDto = append(statisticsDto, dto.MapToStatistics(class))
	}

	return c.JSON(http.StatusOK, statisticsDto)
}

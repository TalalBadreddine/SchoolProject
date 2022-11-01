package controller

import (
	"net/http"
	"server/dto"
	"server/filter"
	"server/repository"

	"github.com/labstack/echo/v4"
)

func GetClasses(c echo.Context) error {

	var classFilter filter.ClassFilter

	err := c.Bind(&classFilter)

	if err != nil {
		return err
	}

	classes := repository.SearchClasses(classFilter)
	var classesDto []*dto.Class

	for _, class := range classes {
		classesDto = append(classesDto, dto.MapClassDto(class))
	}

	return c.JSON(http.StatusOK, classesDto)
}

package controller

import (
	"fmt"
	"net/http"
	"server/dto"
	"server/entity"
	"server/utils"

	"github.com/labstack/echo/v4"
)

func AddClass(c echo.Context) error {
	var class entity.Class = entity.Class{}

	if err := c.Bind(&class); err != nil {
		return err
	}

	message := utils.AddClass(class)

	return c.JSON(http.StatusOK, message)
}

func GetClasses(c echo.Context) error {

	classes := utils.GetAllClasses()
	var classesDto []*dto.Class

	for _, class := range classes {
		classesDto = append(classesDto, dto.MapClassDto(class))
	}

	return c.JSON(http.StatusOK, classesDto)
}

func GetStudentsByClassId(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, id)
}

func GetTeachersByClassId(c echo.Context) error {
	id := c.Param("id")
	results := fmt.Sprintf("Get all Teachers related to classId: %v", id)

	return c.String(http.StatusOK, results)
}

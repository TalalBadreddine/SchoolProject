package controller

import (
	"fmt"
	"net/http"
	"server/dto"
	"server/entity"
	"server/utils"

	"github.com/labstack/echo/v4"
)

func AddTeacher(c echo.Context) error {
	var teacher entity.Teacher

	if err := c.Bind(&teacher); err != nil {
		return err
	}

	results := dto.MapTeacherDto(utils.AddTeacher(teacher))

	return c.JSON(http.StatusOK, results)
}

func AddTeacherToClass(c echo.Context) error {
	var request entity.TeacherClassRequest = entity.TeacherClassRequest{}

	if err := c.Bind(&request); err != nil {
		return err
	}

	message := dto.MapTeacherDto(utils.AddTeacherInClass(request.TeacherId, request.ClassId))

	return c.JSON(http.StatusOK, message)
}

func GetTeachers(c echo.Context) error {
	var filter entity.Filter
	var teachers []*entity.Teacher = utils.SearchTeachers(filter)
	var teachersDto []*dto.Teacher

	for _, teacher := range teachers {
		teachersDto = append(teachersDto, dto.MapTeacherDto(teacher))
	}

	return c.JSON(http.StatusOK, teachersDto)
}

func GetStudentsByTeacherId(c echo.Context) error {
	id := c.Param("id")
	results := fmt.Sprintf("Get all students related to teacher with id = %v", id)
	return c.String(http.StatusOK, results)
}

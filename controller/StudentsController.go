package controller

import (
	"fmt"
	"net/http"
	"server/dto"
	"server/entity"
	"server/utils"

	"github.com/labstack/echo/v4"
)

func AddStudent(c echo.Context) error {
	var student entity.Student = entity.Student{}

	if err := c.Bind(&student); err != nil {
		return err
	}

	results := dto.MapStudentDto(utils.AddStudent(student))

	return c.JSON(http.StatusOK, *results)
}

func AddStudentToClass(c echo.Context) error {
	var request entity.StudentClassRequest = entity.StudentClassRequest{}

	if err := c.Bind(&request); err != nil {
		return err
	}

	message := dto.MapStudentDto(utils.AddStudentInClass(request.StudentId, request.ClassId))

	return c.JSON(http.StatusOK, message)

}

func GetStudents(c echo.Context) error {
	var filter entity.Filter
	var students []*entity.Student = utils.SearchStudents(filter)
	var studentsDto []*dto.Student

	for _, student := range students {
		studentsDto = append(studentsDto, dto.MapStudentDto(student))
	}

	return c.JSON(http.StatusOK, studentsDto)
}

func GetClassesByStudentId(c echo.Context) error {
	id := c.Param("id")
	var filter entity.Filter

	results := utils.SearchStudents(filter)

	fmt.Print(results)

	return c.String(http.StatusOK, id)
}

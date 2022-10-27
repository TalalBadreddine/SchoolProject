package controller

import (
	"fmt"
	"net/http"
	"server/models"

	"github.com/labstack/echo/v4"
)

func AddTeacher(c echo.Context) error {
	var teacher models.Teacher

	if err := c.Bind(&teacher); err != nil {
		return err
	}

	results := models.AddTeacher(teacher)

	return c.String(http.StatusOK, results)
}

func AddTeacherToClass(c echo.Context) error {
	var request models.TeacherClassRequest = models.TeacherClassRequest{}

	if err := c.Bind(&request); err != nil {
		return err
	}

	message := models.AddTeacherInClass(request.TeacherId, request.ClassId)

	return c.String(http.StatusOK, message)
}

func GetTeachers(c echo.Context) error {
	return c.String(http.StatusOK, "Get all teachers")
}

func GetStudentsByTeacherId(c echo.Context) error {
	id := c.Param("id")
	results := fmt.Sprintf("Get all students related to teacher with id = %v", id)
	return c.String(http.StatusOK, results)
}

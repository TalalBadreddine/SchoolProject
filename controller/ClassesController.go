package controller

import (
	"fmt"
	"net/http"
	"server/models"

	"github.com/labstack/echo/v4"
)

func AddClass(c echo.Context) error {
	var class models.Class = models.Class{}

	if err := c.Bind(&class); err != nil {
		return err
	}

	message := models.AddClass(class)

	return c.String(http.StatusOK, message)
}

func GetClasses(c echo.Context) error {
	return c.String(http.StatusOK, "Get all classes")
}

func GetStudentsByClassId(c echo.Context) error {
	id := c.Param("id")
	var filter models.Filter = models.Filter{Class: id}
	var students = models.SearchStudents(filter)

	results := "No Students Are found"

	for _, student := range students {
		results += fmt.Sprintf("The student %v %v With Id: %v is taking this class", student.FirstName, student.LastName, student.ID)
	}

	return c.String(http.StatusOK, results)
}

func GetTeachersByClassId(c echo.Context) error {
	id := c.Param("id")
	results := fmt.Sprintf("Get all Teachers related to classId: %v", id)

	return c.String(http.StatusOK, results)
}

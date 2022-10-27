package controller

import (
	"fmt"
	"net/http"
	"server/models"

	"github.com/labstack/echo/v4"
)

func AddStudent(c echo.Context) error {
	var student models.Student = models.Student{}

	if err := c.Bind(&student); err != nil {
		return err
	}

	results := models.AddStudent(student)

	return c.String(http.StatusOK, fmt.Sprintf(" %v", results))
}

func AddStudentToClass(c echo.Context) error {
	var request models.StudentClassRequest = models.StudentClassRequest{}

	if err := c.Bind(&request); err != nil {
		return err
	}

	message := models.AddStudentInClass(request.StudentId, request.ClassId)

	return c.String(http.StatusOK, message)

}

func GetStudents(c echo.Context) error {
	var students []*models.Student = models.GetAllStudents()
	var results string

	for _, student := range students {
		results += fmt.Sprintf("ID: %v, Student Name: %v, Student Last Name: %v \n", student.ID, student.FirstName, student.LastName)
	}

	return c.String(http.StatusOK, results)
}

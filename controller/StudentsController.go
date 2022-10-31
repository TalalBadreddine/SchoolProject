package controller

import (
	"fmt"
	"net/http"
	"server/dto"
	"server/entity"
	"server/repository"

	"github.com/labstack/echo/v4"
)

func AddStudent(c echo.Context) error {
	var student entity.Student = entity.Student{}

	if err := c.Bind(&student); err != nil {
		return err
	}

	results := dto.MapStudentDto(repository.AddStudent(student))

	return c.JSON(http.StatusOK, *results)
}

func AddStudentToClass(c echo.Context) error {
	var request entity.StudentClassRequest = entity.StudentClassRequest{}

	if err := c.Bind(&request); err != nil {
		return err
	}

	message := dto.MapStudentDto(repository.AddStudentInClass(request.StudentId, request.ClassId))

	return c.JSON(http.StatusOK, message)

}

func GetStudents(c echo.Context) error {
	var filter repository.Filter
	var students []*entity.Student = repository.SearchStudents(filter)
	var studentsDto []*dto.Student

	for _, student := range students {
		studentsDto = append(studentsDto, dto.MapStudentDto(student))
	}

	return c.JSON(http.StatusOK, studentsDto)
}

func GetClassesByStudentId(c echo.Context) error {
	id := c.Param("id")
	var filter repository.Filter

	results := repository.SearchStudents(filter)

	fmt.Print(results)

	return c.String(http.StatusOK, id)
}

func GetStudentById(c echo.Context) error {
	id := c.Param("id")
	var students []dto.Student

	var filter repository.Filter
	fmt.Println(id)

	results := repository.SearchStudents(filter)

	for _, student := range results {
		students = append(students, *dto.MapStudentDto(student))
	}

	return c.JSON(http.StatusOK, students)
}

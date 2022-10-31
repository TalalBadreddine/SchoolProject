package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/dto"
	"server/filter"
	"server/repository"
)

type Student struct {
	studentRepository repository.StudentRepository
}

func ProvideStudent(studentRepository repository.StudentRepository) Student {
	return Student{studentRepository: studentRepository}
}

func (s Student) GetStudents(c echo.Context) error {
	var studentFilter filter.StudentFilter

	err := c.Bind(&studentFilter)

	if err != nil {
		return err
	}

	var students = s.studentRepository.SearchStudents(studentFilter)
	var studentsDto []*dto.Student

	for _, student := range students {
		studentsDto = append(studentsDto, dto.MapStudentDto(student))
	}

	return c.JSON(http.StatusOK, studentsDto)
}

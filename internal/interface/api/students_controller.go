package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/internal/domain/model/filter"
	"server/internal/interface/api/dto"
	"server/internal/useCase/query"
)

type Student struct {
	GetAllStudents query.GetAllStudents
}

func ProvideStudent(s query.GetAllStudents) Student {
	return Student{GetAllStudents: s}
}

func (s Student) GetStudents(c echo.Context) error {
	var studentFilter filter.StudentFilter

	err := c.Bind(&studentFilter)

	if err != nil {
		return err
	}

	var students = s.GetAllStudents.Handle(studentFilter)
	var studentsDto []*dto.Student

	for _, student := range students {
		studentsDto = append(studentsDto, dto.MapStudentDto(student))
	}

	return c.JSON(http.StatusOK, studentsDto)
}

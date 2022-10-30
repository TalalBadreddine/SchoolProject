package controller

import (
	"fmt"
	"net/http"
	"server/dto"
	"server/entity"
	"server/storage"
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
	// id := c.Param("id")
	var db = storage.GetDBInstance()
	var students []entity.Student

	// err := db.Preload("Classes").Joins("JOIN studentsClasses ON studentsClasses.student_id = student.id").
	// 	Joins("JOIN classes ON studentsClasses.class_id = class.id").
	// 	Find(&students)

	err := db.Preload("Classes").Preload("Student").Joins("inner join studentsClasses on studentsClasses.class_id = classes.id and studentsClasses.class_id = 1").
		Find(&students)

	fmt.Print(students)
	fmt.Print(err.Error)

	// var filter entity.Filter = entity.Filter{Class: id}
	// var students = entity.SearchStudents(filter)

	// results := "No Students Are found"

	// for _, student := range students {
	// 	results += fmt.Sprintf("The student %v %v With Id: %v is taking this class", student.FirstName, student.LastName, student.ID)
	// }

	return c.JSON(http.StatusOK, "test")
}

func GetTeachersByClassId(c echo.Context) error {
	id := c.Param("id")
	results := fmt.Sprintf("Get all Teachers related to classId: %v", id)

	return c.String(http.StatusOK, results)
}

// SELECT * FROM studentClasses INNER JOIN student ON student.Id = studentClass.studentId INNER JOIN class ON class.Id = studentClasses.classId WHERE studentClass.classId = %v

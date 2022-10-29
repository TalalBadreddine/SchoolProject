package routes

import (
	"server/controller"

	"github.com/labstack/echo/v4"
)

func InitStudentsRoutes(e *echo.Echo) {
	e.POST("Student", controller.AddStudent)
	e.GET("Student", controller.GetStudents)
	e.PUT("Student/Class", controller.AddStudentToClass)
	e.GET("Student/Class/:id", controller.GetClassesByStudentId)
	// e.GET("Student/Teacher/:id", controller.GetTeachersByStudentsId)
}

// Student/:id/Class -> Join between student, studentClasses, class based od id (where student.id = id)

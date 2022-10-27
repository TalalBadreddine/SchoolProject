package routes

import (
	"server/controller"

	"github.com/labstack/echo/v4"
)

func InitStudentsRoutes(e *echo.Echo) {
	e.POST("Student", controller.AddStudent)
	e.GET("Student", controller.GetStudents)
	e.PUT("Student/Class", controller.AddStudentToClass)
}

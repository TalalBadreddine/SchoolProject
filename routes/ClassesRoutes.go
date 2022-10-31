package routes

import (
	"server/controller"

	"github.com/labstack/echo/v4"
)

func InitClassesRoutes(e *echo.Echo) {
	e.POST("Classes", controller.AddClass)
	e.GET("Classes", controller.GetClasses)
	e.GET("Classes/student/:id", controller.GetStudentsByClassId)
	e.GET("Classes/teacher/:id", controller.GetTeachersByClassId)
}

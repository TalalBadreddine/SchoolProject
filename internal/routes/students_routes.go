package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"server/config"
)

func InitStudentsRoutes(e *echo.Echo, db *gorm.DB) {
	StudentApi := config.WireStudentApi(db)

	e.GET("students", StudentApi.GetAllStudents)
}

package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"server/config"
)

func InitClassesRoutes(e *echo.Echo, db *gorm.DB) {
	ClassApi := config.WireClassApi(db)

	e.GET("classes", ClassApi.GetClasses)
}

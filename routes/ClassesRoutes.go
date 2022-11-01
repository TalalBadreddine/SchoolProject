package routes

import (
	"server/controller"

	"github.com/labstack/echo/v4"
)

func InitClassesRoutes(e *echo.Echo) {

	e.GET("Classes", controller.GetClasses)
}

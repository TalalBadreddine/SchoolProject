package main

import (
	"server/models"
	"server/routes"
	"server/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	storage.NewDB()

	// storage.GetDBInstance().DropTableIfExists(&models.Student{}, &models.Class{}, &models.Teacher{})
	storage.GetDBInstance().AutoMigrate(&models.Student{}, &models.Class{}, &models.Teacher{})

	routes.InitStudentsRoutes(e)
	routes.InitTeachersRoutes(e)
	routes.InitClassesRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}

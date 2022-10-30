package main

import (
	"server/entity"
	"server/routes"
	"server/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	storage.NewDB()

	// storage.GetDBInstance().DropTableIfExists(&entity.Student{}, &entity.Class{}, &entity.Teacher{})
	storage.GetDBInstance().AutoMigrate(&entity.Student{}, &entity.Class{}, &entity.Teacher{})

	routes.InitStudentsRoutes(e)
	routes.InitTeachersRoutes(e)
	routes.InitClassesRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}

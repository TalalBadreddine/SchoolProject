package main

import (
	"server/entity"
	"server/routes"
	"server/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := storage.NewDB()

	err := db.AutoMigrate(&entity.Student{}, &entity.Class{}, &entity.Teacher{})
	if err != nil {
		return
	}

	routes.InitStudentsRoutes(e, db)
	routes.InitClassesRoutes(e, db)

	//TODO move port to a variable stored in config file
	e.Logger.Fatal(e.Start(":1323"))
}

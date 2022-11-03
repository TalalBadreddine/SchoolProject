package main

import (
	"server/config"
	"server/internal/adapter/db/entity"
	"server/internal/routes"
	"server/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := storage.NewDB()

	err := db.AutoMigrate(&entity.Student{}, &entity.Class{}, &entity.Teacher{}, &entity.StudentsClasses{})
	if err != nil {
		return
	}

	routes.InitStudentsRoutes(e, db)
	routes.InitClassesRoutes(e, db)

	e.Logger.Fatal(e.Start(config.ConfigInstance.SERVERPORT))
}

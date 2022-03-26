package main

import (
	"cake-store-api/config"
	"cake-store-api/controller"
	"cake-store-api/database"
	"cake-store-api/repository"
	"cake-store-api/service"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := database.InitMySQL()
	if err != nil {
		panic(err)
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	controller := controller.NewController(service)

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: config.LogFormat + "\n",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/cakes", controller.GetCakeList)
	e.GET("/cakes/:cake_id", controller.GetCake)

	e.POST("/cakes", controller.CreateCake)
	e.PATCH("/cakes/:cake_id", controller.UpdateCake)
	e.DELETE("/cakes/:cake_id", controller.DeleteCake)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

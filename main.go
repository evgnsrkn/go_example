package main

import (
	"go_example/db"
	"go_example/service"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	db.Init()
	e := echo.New()

	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Routes
	e.GET("/users", service.GetAll)
	e.GET("/users/:id", service.GetById)
	e.POST("/users", service.Create)
	e.DELETE("/users/:id", service.Delete)
	e.PUT("/users/:id", service.Update)

	if err := e.Start(":7070"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

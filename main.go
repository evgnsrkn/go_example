package main

import (
	"go_example/db"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	db.ConnectDatabase()

	e.Logger.Fatal(e.Start(":1323"))
}

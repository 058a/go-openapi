package main

import (
	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	hello "openapi/internal/hello/api"
	stockitem "openapi/internal/stockitem/api"
)

// main is the entry point of the program.
//
// No parameters.
// No return type.
func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello.Get)
	e.POST("/stock/items", stockitem.Post)
	e.PUT("/stock/items/:id", stockitem.Put)

	e.Logger.Fatal(e.Start(":3000"))
}

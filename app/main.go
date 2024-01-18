package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":3000"))
}

type ResponseData struct {
	Message string `json:"message"`
}

func hello(c echo.Context) error {
	data := ResponseData{
		Message: "Hello, World!!",
	}
	return c.JSON(http.StatusOK, data)
}

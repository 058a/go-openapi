package main

import (
	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"openapi/internal/infra/oapi_codegen/hello_api"
	"openapi/internal/infra/oapi_codegen/stockitem_api"

	oapiMiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"

	"openapi/internal/hello"
	"openapi/internal/stockitem"
)

// main is the entry point of the program.
//
// No parameters.
// No return type.
func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	helloSwagger, err := hello_api.GetSwagger()
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.Use(oapiMiddleware.OapiRequestValidator(helloSwagger))

	stockitemSwagger, err := stockitem_api.GetSwagger()
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.Use(oapiMiddleware.OapiRequestValidator(stockitemSwagger))

	e.GET("/", hello.GetHello)
	e.POST("/stock/items", stockitem.PostStockItem)

	e.Logger.Fatal(e.Start(":3000"))
}

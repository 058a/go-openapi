package main

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"openapi/application/hello"
	stock_item "openapi/application/stock_item"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	dbDriver := "postgres"
	dsn := "host=openapi-db port=5432 user=user password=password dbname=openapi sslmode=disable"

	db, openErr := sql.Open(dbDriver, dsn)
	if openErr != nil {
		e.Logger.Fatal(openErr)
	}
	defer db.Close()

	e.GET("/", hello.GetHello)
	e.POST("/stock/items", func(ctx echo.Context) error {
		return stock_item.PostStockItem(ctx, db)
	})
	e.Logger.Fatal(e.Start(":3000"))
}

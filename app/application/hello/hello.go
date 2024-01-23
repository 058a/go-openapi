package hello

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetHello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &Hello{
		Message: "Hello, World!",
	})
}

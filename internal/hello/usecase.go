package hello

import (
	"net/http"
	"openapi/internal/infra/oapi_codegen/hello_api"

	"github.com/labstack/echo/v4"
)

func GetHello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &hello_api.Hello{
		Message: "Hello, World!",
	})
}

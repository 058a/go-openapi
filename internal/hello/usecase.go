package hello

import (
	"net/http"
	"openapi/internal/infra/oapi_codegen/hello_api"

	"github.com/labstack/echo/v4"
)

// GetHello is a function that returns a JSON response with a message "Hello, World!".
//
// It takes in a parameter of type echo.Context and returns an error.
func GetHello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &hello_api.Hello{
		Message: "Hello, World!",
	})
}

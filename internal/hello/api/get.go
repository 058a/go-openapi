package hello

import (
	"net/http"
	openapi "openapi/internal/infra/oapi_codegen/hello"

	"github.com/labstack/echo/v4"
)

// GetHello is a function that returns a JSON response with a message "Hello, World!".
//
// It takes in a parameter of type echo.Context and returns an error.
func Get(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, &openapi.Hello{
		Message: "Hello, World!",
	})
}

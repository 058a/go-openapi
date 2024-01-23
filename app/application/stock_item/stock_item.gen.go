// Package stock_item provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package stock_item

import (
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// BadRequestResponse defines model for BadRequestResponse.
type BadRequestResponse = []struct {
	Message *string `json:"message,omitempty"`
}

// CreatedResponse defines model for CreatedResponse.
type CreatedResponse struct {
	Id openapi_types.UUID `json:"id"`
}

// BadRequest defines model for BadRequest.
type BadRequest = BadRequestResponse

// Created defines model for Created.
type Created = CreatedResponse

// PostStockItemJSONBody defines parameters for PostStockItem.
type PostStockItemJSONBody struct {
	Name string `json:"name"`
}

// PostStockItemJSONRequestBody defines body for PostStockItem for application/json ContentType.
type PostStockItemJSONRequestBody PostStockItemJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create Stock Item
	// (POST /stock/items)
	PostStockItem(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostStockItem converts echo context to params.
func (w *ServerInterfaceWrapper) PostStockItem(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostStockItem(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/stock/items", wrapper.PostStockItem)

}

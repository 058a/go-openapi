package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"openapi/internal/infra/database"
	openapi "openapi/internal/infra/oapi_codegen/stockitem"

	"openapi/internal/stockitem/usecase"
)

// PostStockItem is a function that handles the HTTP POST request for creating a new stock item.
func Post(ctx echo.Context) error {
	request := &openapi.PostStockItemJSONBody{}
	ctx.Bind(&request)

	UnverifiedRequestDto := usecase.UnverifiedCreateRequestDto{Name: request.Name}
	verifiedRequestDto, verfyErr := UnverifiedRequestDto.Verify()
	if verfyErr != nil {
		return ctx.JSON(http.StatusBadRequest, verfyErr)
	}

	db, dbErr := database.New()
	if dbErr != nil {
		return ctx.JSON(http.StatusInternalServerError, dbErr)
	}
	defer db.Close()

	responseDto, err := usecase.CreateStockItemUseCase(verifiedRequestDto, db)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	response := &openapi.Created{
		Id: responseDto.Id,
	}

	return ctx.JSON(http.StatusCreated, response)
}

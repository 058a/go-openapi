package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"openapi/internal/infra/database"
	"openapi/internal/infra/oapi_codegen/stockitem_api"

	"openapi/internal/stockitem/usecase"
)

// PostStockItem is a function that handles the HTTP POST request for creating a new stock item.
//
// It takes in a context object and returns an error.
func PostStockItem(ctx echo.Context) error {
	request := &stockitem_api.PostStockItemJSONBody{}
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

	response := &stockitem_api.Created{
		Id: responseDto.Id,
	}

	return ctx.JSON(http.StatusCreated, response)
}

func PutStockItem(ctx echo.Context) error {
	request := &stockitem_api.PutStockItemJSONRequestBody{}
	ctx.Bind(&request)

	unverifiedRequestDto := usecase.UnverifiedUpdateRequestDto{Name: request.Name}
	verifiedRequestDto, verfyErr := unverifiedRequestDto.Verify()
	if verfyErr != nil {
		return ctx.JSON(http.StatusBadRequest, verfyErr)
	}

	db, dbErr := database.New()
	if dbErr != nil {
		return ctx.JSON(http.StatusInternalServerError, dbErr)
	}
	defer db.Close()

	responseDto, err := usecase.UpdateStockItemUseCase(*verifiedRequestDto, db)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	response := &stockitem_api.Created{
		Id: responseDto.Id,
	}

	return ctx.JSON(http.StatusOK, response)
}

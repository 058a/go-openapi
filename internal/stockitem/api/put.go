package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"openapi/internal/infra/database"
	oapicodegen "openapi/internal/infra/oapi_codegen/stockitem"

	"openapi/internal/stockitem/usecase"
)

// PutStockItem is a function that handles the HTTP PUT request for updating an existing stock item.
func Put(ctx echo.Context) error {
	request := &oapicodegen.PutStockItemJSONRequestBody{}
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

	response := &oapicodegen.Created{
		Id: responseDto.Id,
	}

	return ctx.JSON(http.StatusOK, response)
}

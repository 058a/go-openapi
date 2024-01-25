package stockitem

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"openapi/internal/infra/database"
	"openapi/internal/infra/oapi_codegen/stockitem_api"
)

// PostStockItem is a function that handles the HTTP POST request for creating a new stock item.
//
// It takes in a context object and returns an error.
func PostStockItem(ctx echo.Context) error {
	request := &stockitem_api.PostStockItemJSONBody{}
	ctx.Bind(&request)

	UnverifiedRequestDto := UnverifiedCreateUseCaseRequestDto{request.Name}
	verifiedRequestDto, verfyErr := UnverifiedRequestDto.Verify()
	if verfyErr != nil {
		return ctx.JSON(http.StatusBadRequest, verfyErr)
	}

	db, dbErr := database.New()
	if dbErr != nil {
		return ctx.JSON(http.StatusInternalServerError, dbErr)
	}
	defer db.Close()
	repository := &StockItemRepository{db}

	responseDto, err := CreateStockItemUseCase(*verifiedRequestDto, repository)
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

	unverifiedRequestDto := UnverifiedUpdateUseCaseRequestDto{request.Name}
	VerifiedRequestDto, verfyErr := unverifiedRequestDto.Verify()
	if verfyErr != nil {
		return ctx.JSON(http.StatusBadRequest, verfyErr)
	}

	db, dbErr := database.New()
	if dbErr != nil {
		return ctx.JSON(http.StatusInternalServerError, dbErr)
	}
	defer db.Close()
	repository := &StockItemRepository{db}

	responseDto, err := UpdateStockItemUseCase(*VerifiedRequestDto, repository)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	response := &stockitem_api.Created{
		Id: responseDto.Id,
	}

	return ctx.JSON(http.StatusOK, response)
}

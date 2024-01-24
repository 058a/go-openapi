package stockitem

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"openapi/internal/infra/database"
	"openapi/internal/infra/oapi_codegen/stockitem_api"
)

func PostStockItem(ctx echo.Context) error {
	request := &stockitem_api.PostStockItemJSONBody{}
	ctx.Bind(&request)

	requestDto := PostStockItemUseCaseRequestDto{request.Name}

	db, dbErr := database.New()
	if dbErr != nil {
		return ctx.JSON(http.StatusInternalServerError, dbErr)
	}
	defer db.Close()
	repository := &StockItemRepository{db}

	responseDto, err := PostStockItemUseCase(requestDto, repository)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	response := &stockitem_api.Created{
		Id: responseDto.Id,
	}

	return ctx.JSON(http.StatusCreated, response)
}

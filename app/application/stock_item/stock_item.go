package stock_item

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	domain "openapi/domain/models"
	"openapi/repository"
)

func PostStockItem(ctx echo.Context, db *sql.DB) error {
	request := &PostStockItemJSONBody{}
	ctx.Bind(&request)

	stockItem := domain.NewStockItem(request.Name)

	stockItemRepository := repository.StockItemRepository{}
	storeErr := stockItemRepository.Save(db, *stockItem)
	if storeErr != nil {
		return ctx.JSON(http.StatusInternalServerError, storeErr)
	}

	createdResponse := &StockItem{
		Id: stockItem.Id,
	}

	return ctx.JSON(http.StatusCreated, createdResponse)
}

func PutStockItem(ctx echo.Context, db *sql.DB) error {
	request := &PostStockItemJSONBody{}
	ctx.Bind(&request)
	stockItemId := ctx.Param("id")

	return ctx.JSON(http.StatusOK, nil)
}

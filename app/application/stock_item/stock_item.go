package stock_item

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	domain "openapi/domain/models"
	"openapi/repository"
)

func PostStockItem(ctx echo.Context, db *sql.DB) error {
	request := &PostStockItemJSONBody{}
	ctx.Bind(&request)

	stockItem := domain.StockItem{
		Id:   uuid.New(),
		Name: request.Name,
	}

	stockItemRepository := repository.StockItem{}
	storeErr := stockItemRepository.Insert(db, stockItem)
	if storeErr != nil {
		return ctx.JSON(http.StatusInternalServerError, storeErr)
	}

	createdResponse := &CreatedResponse{
		Id: stockItem.Id,
	}

	return ctx.JSON(http.StatusCreated, createdResponse)
}
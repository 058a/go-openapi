package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"openapi/internal/infra/database"
	oapicodegen "openapi/internal/infra/oapicodegen/stockitem"

	"openapi/internal/stockitem/usecase"
)

// PostStockItem is a function that handles the HTTP POST request for creating a new stock item.
func Post(c echo.Context) error {
	request := &oapicodegen.PostStockItemJSONBody{}
	c.Bind(&request)

	UnverifiedRequestDto := usecase.UnverifiedCreateRequestDto{Name: request.Name}
	verifiedRequestDto, verfyErr := UnverifiedRequestDto.Verify()
	if verfyErr != nil {
		return c.JSON(http.StatusBadRequest, verfyErr)
	}

	db, dbErr := database.New()
	if dbErr != nil {
		return c.JSON(http.StatusInternalServerError, dbErr)
	}
	defer db.Close()

	responseDto, err := usecase.CreateStockItemUseCase(verifiedRequestDto, db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response := &oapicodegen.Created{
		Id: responseDto.Id,
	}

	return c.JSON(http.StatusCreated, response)
}

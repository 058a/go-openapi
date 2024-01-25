package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"openapi/internal/infra/database"
	oapicodegen "openapi/internal/infra/oapi_codegen/stockitem"

	"openapi/internal/stockitem/repository"
	"openapi/internal/stockitem/usecase"
)

// PutStockItem is a function that handles the HTTP PUT request for updating an existing stock item.
func Put(c echo.Context) error {
	id := c.Param("id")

	request := &oapicodegen.PutStockItemJSONRequestBody{}
	c.Bind(&request)

	unverifiedRequestDto := usecase.UnverifiedUpdateRequestDto{
		Id:   id,
		Name: request.Name}

	db, dbErr := database.New()
	if dbErr != nil {
		return c.JSON(http.StatusInternalServerError, dbErr)
	}
	defer db.Close()

	verifiedRequestDto, verfyErr := unverifiedRequestDto.Verify()
	if verfyErr != nil {
		return c.JSON(http.StatusBadRequest, verfyErr)
	}

	_, getErr := repository.Get(db, verifiedRequestDto.Id)
	if getErr != nil {
		return c.JSON(http.StatusNotFound, getErr)
	}

	responseDto, err := usecase.UpdateStockItemUseCase(*verifiedRequestDto, db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response := &oapicodegen.Created{
		Id: responseDto.Id,
	}

	return c.JSON(http.StatusOK, response)
}

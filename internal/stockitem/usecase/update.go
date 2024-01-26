package usecase

import (
	"database/sql"
	"fmt"
	"openapi/internal/stockitem/model"
	"openapi/internal/stockitem/repository"
	"strings"

	"github.com/google/uuid"
)

type UnverifiedUpdateRequestDto struct {
	Id   string
	Name string
}

func (s UnverifiedUpdateRequestDto) Verify() (*VerifiedUpdateRequestDto, error) {
	trimedName := strings.TrimSpace(s.Name)
	if !(len(trimedName) > 0 && len(trimedName) <= 100) {
		return nil, fmt.Errorf("name length should be between 1 and 100")
	}

	parcedId := uuid.MustParse(s.Id)

	return &VerifiedUpdateRequestDto{
		Id:   parcedId,
		Name: s.Name,
	}, nil
}

type VerifiedUpdateRequestDto struct {
	Id   uuid.UUID
	Name string
}

func (s VerifiedUpdateRequestDto) GenerateModel() *model.StockItemModel {
	return model.Renew(s.Id, s.Name)
}

type UpdateResponseDto struct {
	Id uuid.UUID
}

// UpdateStockItemUseCase is a Go function that handles updating a stock item.
func UpdateStockItemUseCase(
	requestDto *VerifiedUpdateRequestDto,
	db *sql.DB) (UpdateResponseDto, error) {

	stockItem := requestDto.GenerateModel()

	saveErr := repository.Save(db, *stockItem)
	if saveErr != nil {
		return UpdateResponseDto{}, saveErr
	}

	return UpdateResponseDto{stockItem.Id}, nil
}
package usecase

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"

	"openapi/internal/stockitem/model"
	"openapi/internal/stockitem/repository"
)

type UnverifiedCreateRequestDto struct {
	Name string
}

func (s UnverifiedCreateRequestDto) Verify() (VerifiedCreateRequestDto, error) {
	trimedName := strings.TrimSpace(s.Name)
	if !(len(trimedName) > 0 && len(trimedName) <= 100) {
		return VerifiedCreateRequestDto{}, fmt.Errorf("name length should be between 1 and 100")
	}

	return VerifiedCreateRequestDto{s.Name}, nil
}

type VerifiedCreateRequestDto struct {
	Name string
}

func (s VerifiedCreateRequestDto) GenerateModel() *model.StockItemModel {
	return model.New(s.Name)
}

type CreateResponseDto struct {
	Id uuid.UUID
}

// PostStockItemUseCase is a Go function that handles posting a stock item.
//
// It takes a requestDto of type PostStockItemUseCaseRequestDto and a repository of type *StockItemRepository as parameters.
// It returns a PostStockItemUseCaseResponseDto and an error.
func CreateStockItemUseCase(
	requestDto VerifiedCreateRequestDto,
	db *sql.DB) (CreateResponseDto, error) {

	stockItem := requestDto.GenerateModel()

	saveErr := repository.Save(db, *stockItem)
	if saveErr != nil {
		return CreateResponseDto{}, saveErr
	}

	return CreateResponseDto{stockItem.Id}, nil
}

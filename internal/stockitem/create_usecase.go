package stockitem

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type UnverifiedCreateUseCaseRequestDto struct {
	Name string
}

func (s UnverifiedCreateUseCaseRequestDto) Verify() (*VerifiedCreateUseCaseRequestDto, error) {
	trimedName := strings.TrimSpace(s.Name)
	if !(len(trimedName) > 0 && len(trimedName) <= 100) {
		return nil, fmt.Errorf("name length should be between 1 and 100")
	}

	return &VerifiedCreateUseCaseRequestDto{s.Name}, nil
}

type VerifiedCreateUseCaseRequestDto struct {
	Name string
}

func (s VerifiedCreateUseCaseRequestDto) GenerateModel() *StockItemModel {
	return NewStockItemModel(s.Name)
}

type CreateUseCaseResponseDto struct {
	Id uuid.UUID
}

// PostStockItemUseCase is a Go function that handles posting a stock item.
//
// It takes a requestDto of type PostStockItemUseCaseRequestDto and a repository of type *StockItemRepository as parameters.
// It returns a PostStockItemUseCaseResponseDto and an error.
func CreateStockItemUseCase(
	requestDto VerifiedCreateUseCaseRequestDto,
	repository *StockItemRepository) (CreateUseCaseResponseDto, error) {

	stockItem := requestDto.GenerateModel()

	saveErr := repository.Save(*stockItem)
	if saveErr != nil {
		return CreateUseCaseResponseDto{}, saveErr
	}

	return CreateUseCaseResponseDto{stockItem.Id}, nil
}

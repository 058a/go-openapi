package stockitem

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type UnverifiedUpdateUseCaseRequestDto struct {
	Name string
}

func (s UnverifiedUpdateUseCaseRequestDto) Verify() (*VerifiedUpdateUseCaseRequestDto, error) {
	trimedName := strings.TrimSpace(s.Name)
	if !(len(trimedName) > 0 && len(trimedName) <= 100) {
		return nil, fmt.Errorf("name length should be between 1 and 100")
	}
	return &VerifiedUpdateUseCaseRequestDto{s.Name}, nil
}

type VerifiedUpdateUseCaseRequestDto struct {
	Name string
}

func (s VerifiedUpdateUseCaseRequestDto) GenerateModel() *StockItemModel {
	return NewStockItemModel(s.Name)
}

type UpdateUseCaseResponseDto struct {
	Id uuid.UUID
}

// UpdateStockItemUseCase is a Go function that handles updating a stock item.
func UpdateStockItemUseCase(
	requestDto VerifiedUpdateUseCaseRequestDto,
	repository *StockItemRepository) (UpdateUseCaseResponseDto, error) {

	stockItem := requestDto.GenerateModel()

	saveErr := repository.Save(*stockItem)
	if saveErr != nil {
		return UpdateUseCaseResponseDto{}, saveErr
	}

	return UpdateUseCaseResponseDto{stockItem.Id}, nil
}

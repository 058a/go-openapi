package stockitem

import (
	"github.com/google/uuid"
)

type PostStockItemUseCaseRequestDto struct {
	Name string
}
type PostStockItemUseCaseResponseDto struct {
	Id uuid.UUID
}

// PostStockItemUseCase is a Go function that handles posting a stock item.
//
// It takes a requestDto of type PostStockItemUseCaseRequestDto and a repository of type *StockItemRepository as parameters.
// It returns a PostStockItemUseCaseResponseDto and an error.
func PostStockItemUseCase(
	requestDto PostStockItemUseCaseRequestDto,
	repository *StockItemRepository) (PostStockItemUseCaseResponseDto, error) {

	stockItem := NewStockItemModel(requestDto.Name)

	saveErr := repository.Save(*stockItem)
	if saveErr != nil {
		return PostStockItemUseCaseResponseDto{}, saveErr
	}

	return PostStockItemUseCaseResponseDto{stockItem.Id}, nil
}

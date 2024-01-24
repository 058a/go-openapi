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

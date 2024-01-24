package stockitem

import (
	"openapi/internal/infra/database"
	"testing"

	"github.com/google/uuid"
)

func TestPostStockItemUseCase(t *testing.T) {

	requestDto := PostStockItemUseCaseRequestDto{uuid.NewString()}

	db, dbErr := database.New()
	if dbErr != nil {
		t.Fatal(dbErr)
	}
	defer db.Close()
	repository := &StockItemRepository{db}

	responseDto, err := PostStockItemUseCase(requestDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	if responseDto.Id == uuid.Nil {
		t.Errorf("want not nil, got nil")
	}
}

package stockitem

import (
	"openapi/internal/infra/database"
	"testing"

	"github.com/google/uuid"
)

// TestPostStockItemUseCase is a test function for the PostStockItemUseCase.
// It tests the functionality of the PostStockItemUseCase with various scenarios.
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

	model, getErr := repository.Get(responseDto.Id)
	if getErr != nil {
		t.Fatal(getErr)
	}

	if model.Name != requestDto.Name {
		t.Errorf("want %s, got %s", requestDto.Name, model.Name)
	}
}

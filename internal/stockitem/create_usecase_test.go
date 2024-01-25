package stockitem

import (
	"openapi/internal/infra/database"
	"testing"

	"github.com/google/uuid"
)

// TestCreateStockItemUseCase is a test function for the CreateStockItemUseCase.
// It tests the functionality of the CreateStockItemUseCase with various scenarios.
func TestCreateUseCase(t *testing.T) {
	requestDto := VerifiedCreateUseCaseRequestDto{uuid.NewString()}

	db, dbErr := database.New()
	if dbErr != nil {
		t.Fatal(dbErr)
	}
	defer db.Close()
	repository := &StockItemRepository{db}

	responseDto, err := CreateStockItemUseCase(requestDto, repository)
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

	if string(model.Name) != requestDto.Name {
		t.Errorf("want %s, got %s", requestDto.Name, model.Name)
	}
}

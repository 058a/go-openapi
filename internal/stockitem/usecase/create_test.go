package usecase

import (
	"openapi/internal/infra/database"
	"testing"

	"github.com/google/uuid"

	"openapi/internal/stockitem/repository"
)

// TestCreateStockItemUseCase is a test function for the CreateStockItemUseCase.
// It tests the functionality of the CreateStockItemUseCase with various scenarios.
func TestCreateUseCase(t *testing.T) {
	requestDto := VerifiedCreateRequestDto{uuid.NewString()}

	db, dbErr := database.New()
	if dbErr != nil {
		t.Fatal(dbErr)
	}
	defer db.Close()

	responseDto, err := CreateStockItemUseCase(requestDto, db)
	if err != nil {
		t.Fatal(err)
	}

	if responseDto.Id == uuid.Nil {
		t.Errorf("want not nil, got nil")
	}

	model, getErr := repository.Get(db, responseDto.Id)
	if getErr != nil {
		t.Fatal(getErr)
	}

	if string(model.Name) != requestDto.Name {
		t.Errorf("want %s, got %s", requestDto.Name, model.Name)
	}
}

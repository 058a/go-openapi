package usecase

import (
	"openapi/internal/infra/database"
	"testing"

	"github.com/google/uuid"

	"openapi/internal/stockitem/repository"
)

// TestCreateStockItemUseCase is a test function for the CreateStockItemUseCase.
// It tests the functionality of the CreateStockItemUseCase with various scenarios.
func TestCreate(t *testing.T) {

	UnverifiedRequestDto := UnverifiedCreateRequestDto{
		Name: uuid.NewString(),
	}
	verifiedRequestDto, verfyErr := UnverifiedRequestDto.Verify()
	if verfyErr != nil {
		t.Fatal(verfyErr)
	}

	db, dbErr := database.New()
	if dbErr != nil {
		t.Fatal(dbErr)
	}
	defer db.Close()

	responseDto, err := CreateStockItemUseCase(verifiedRequestDto, db)
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

	if string(model.Name) != verifiedRequestDto.Name {
		t.Errorf("want %s, got %s", verifiedRequestDto.Name, model.Name)
	}
}

package usecase

import (
	"testing"

	"openapi/internal/infra/database"
	"openapi/internal/stockitem/repository"

	"github.com/google/uuid"
)

func TestUpdateUseCase(t *testing.T) {

	UnverifiedRequestDto := UnverifiedUpdateRequestDto{
		Id:   uuid.NewString(),
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

	responseDto, err := UpdateStockItemUseCase(verifiedRequestDto, db)
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

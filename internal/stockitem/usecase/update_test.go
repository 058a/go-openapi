package usecase

import (
	"testing"

	"openapi/internal/infra/database"
	"openapi/internal/stockitem/repository"

	"github.com/google/uuid"
)

func TestUpdateUseCase(t *testing.T) {

	requestDto := VerifiedUpdateRequestDto{
		Id:   uuid.New(),
		Name: uuid.NewString()}

	db, dbErr := database.New()
	if dbErr != nil {
		t.Fatal(dbErr)
	}
	defer db.Close()

	responseDto, err := UpdateStockItemUseCase(requestDto, db)
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

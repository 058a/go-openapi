package repository

import (
	"testing"

	_ "github.com/lib/pq"

	"openapi/internal/infra/database"

	"github.com/google/uuid"

	"openapi/internal/stockitem/model"
)

// TestStockItemRepository is a Go function for testing the StockItemRepository.
//
// It takes a testing.T parameter and does not return anything.
func TestStockItemRepository(t *testing.T) {

	db, dbErr := database.New()
	if dbErr != nil {
		t.Fatal(dbErr)
	}
	defer db.Close()

	stockItem := model.New(uuid.NewString())

	storeErr := Save(db, *stockItem)
	if storeErr != nil {
		t.Fatal(storeErr)
	}

	getStockItem, getErr := Get(db, stockItem.Id)
	if getErr != nil {
		t.Fatal(getErr)
	}

	if getStockItem.Id != stockItem.Id {
		t.Errorf("want %s, got %s", stockItem.Id, getStockItem.Id)
	}

	if getStockItem.Name != stockItem.Name {
		t.Errorf("want %s, got %s", stockItem.Name, getStockItem.Name)
	}
}

package stockitem

import (
	"testing"

	_ "github.com/lib/pq"

	"openapi/internal/infra/database"

	"github.com/google/uuid"
)

func TestStockItemRepository(t *testing.T) {

	db, dbErr := database.New()
	if dbErr != nil {
		t.Fatal(dbErr)
	}
	defer db.Close()

	stockItem := NewStockItemModel(uuid.NewString())

	stockItemRepository := &StockItemRepository{db}
	storeErr := stockItemRepository.Save(*stockItem)
	if storeErr != nil {
		t.Fatal(storeErr)
	}

	getStockItem, getErr := stockItemRepository.Get(stockItem.Id)
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

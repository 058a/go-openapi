package stockitem

import (
	"testing"

	"github.com/google/uuid"
)

func TestStockItem(t *testing.T) {

	stockItem := NewStockItemModel("test")
	if stockItem.Id == uuid.Nil {
		t.Errorf("expected not empty, actual empty")
	}

	if stockItem.Name != "test" {
		t.Errorf("expected %s, actual %s", "test", stockItem.Name)
	}
}

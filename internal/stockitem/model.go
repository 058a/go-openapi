package stockitem

import (
	"github.com/google/uuid"
)

type StockItemModel struct {
	Id   uuid.UUID
	Name string
}

// NewStockItemModel creates a new StockItemModel.
//
// It takes a name string as a parameter and returns a pointer to StockItemModel.
func NewStockItemModel(name string) *StockItemModel {
	return &StockItemModel{
		Id:   uuid.New(),
		Name: name,
	}
}

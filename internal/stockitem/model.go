package stockitem

import (
	"github.com/google/uuid"
)

type StockItemModel struct {
	Id   uuid.UUID
	Name string
}

func NewStockItemModel(name string) *StockItemModel {
	return &StockItemModel{
		Id:   uuid.New(),
		Name: name,
	}
}

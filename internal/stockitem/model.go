package stockitem

import (
	"github.com/google/uuid"
)

type StockItemModel struct {
	Id   uuid.UUID
	Name string
}

func NewStockItem(Name string) *StockItemModel {
	return &StockItemModel{
		Id:   uuid.New(),
		Name: Name,
	}
}

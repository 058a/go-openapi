package domain

import (
	"github.com/google/uuid"
)

type StomItem struct {
	Id   uuid.UUID
	Name string
}

func NewStockItem(Name string) *StockItem {
	return &StockItem{
		Id:   uuid.New(),
		Name: Name,
	}
}

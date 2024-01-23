package domain

import "github.com/google/uuid"

type StockItem struct {
	Id   uuid.UUID
	Name string
}

package repository

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"

	domain "openapi/domain/models"
	repository "openapi/repository/models"
)

type StockItem struct{}

func (s *StockItem) Insert(db *sql.DB, stockItem domain.StockItem) error {

	// dbPingErr := db.Ping()
	// if dbPingErr != nil {
	// 	return dbPingErr
	// }

	StockItem := repository.StockItem{
		ID:   stockItem.Id.String(),
		Name: stockItem.Name,
	}

	execErr := StockItem.Insert(context.Background(), db, boil.Infer())
	if execErr != nil {
		return execErr
	}

	return nil
}

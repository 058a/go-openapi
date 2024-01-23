package repository

import (
	"context"
	"database/sql"

	domain "openapi/domain/models"
	repository "openapi/repository/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type StockItem struct{}

func (s *StockItem) Insert(db *sql.DB, stockItem domain.StockItem) error {

	stockItemRecord := repository.StockItem{
		ID:   stockItem.Id.String(),
		Name: stockItem.Name,
	}

	dbExecErr := stockItemRecord.Insert(context.Background(), db, boil.Infer())
	if dbExecErr != nil {
		return dbExecErr
	}

	return nil
}

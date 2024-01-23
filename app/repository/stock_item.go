package repository

import (
	"context"
	"database/sql"

	domain "openapi/domain/models"
	repository_models "openapi/repository/models"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type StockItemRepository struct{}

func (r *StockItemRepository) Save(db *sql.DB, stockItem domain.StockItem) error {

	exists, findErr := repository_models.FindStockItem(context.Background(), db, stockItem.Id.String())
	stockItemRecord := repository_models.StockItem{}
	if findErr != nil {
		stockItemRecord.ID = stockItem.Id.String()
		stockItemRecord.Name = stockItem.Name
		dbExecErr := stockItemRecord.Insert(context.Background(), db, boil.Infer())
		if dbExecErr != nil {
			return dbExecErr
		}
	} else {
		stockItemRecord.ID = exists.ID
		stockItemRecord.Name = stockItem.Name
		stockItemRecord.Update(context.Background(), db, boil.Infer())
	}

	return nil
}

func (r *StockItemRepository) Get(db *sql.DB, id uuid.UUID) (domain.StockItem, error) {
	exists, findErr := repository_models.FindStockItem(context.Background(), db, id.String())
	if findErr != nil {
		return domain.StockItem{}, findErr
	}

	stockItem := domain.StockItem{
		Id:   uuid.MustParse(exists.ID),
		Name: exists.Name,
	}
	return stockItem, nil
}

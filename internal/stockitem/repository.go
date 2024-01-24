package stockitem

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"openapi/internal/infra/sqlboiler"
)

type StockItemRepository struct {
	db *sql.DB
}

func (r *StockItemRepository) Save(stockItem StockItemModel) error {

	exists, findErr := sqlboiler.FindStockItem(context.Background(), r.db, stockItem.Id.String())
	stockItemRecord := sqlboiler.StockItem{}
	if findErr != nil {
		stockItemRecord.ID = stockItem.Id.String()
		stockItemRecord.Name = stockItem.Name
		dbExecErr := stockItemRecord.Insert(context.Background(), r.db, boil.Infer())
		if dbExecErr != nil {
			return dbExecErr
		}
	} else {
		stockItemRecord.ID = exists.ID
		stockItemRecord.Name = stockItem.Name
		stockItemRecord.Update(context.Background(), r.db, boil.Infer())
	}

	return nil
}

func (r *StockItemRepository) Get(id uuid.UUID) (StockItemModel, error) {
	exists, findErr := sqlboiler.FindStockItem(context.Background(), r.db, id.String())
	if findErr != nil {
		return StockItemModel{}, findErr
	}

	stockItem := StockItemModel{
		Id:   uuid.MustParse(exists.ID),
		Name: exists.Name,
	}
	return stockItem, nil
}

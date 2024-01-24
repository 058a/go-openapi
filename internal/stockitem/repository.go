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

// Save saves the stock item model to the repository.
// It takes a StockItemModel as a parameter and returns an error.
func (r *StockItemRepository) Save(model StockItemModel) error {

	exists, findErr := sqlboiler.FindStockItem(context.Background(), r.db, model.Id.String())
	stockItem := sqlboiler.StockItem{}
	if findErr != nil {
		stockItem.ID = model.Id.String()
		stockItem.Name = model.Name
		dbExecErr := stockItem.Insert(context.Background(), r.db, boil.Infer())
		if dbExecErr != nil {
			return dbExecErr
		}
	} else {
		stockItem.ID = exists.ID
		stockItem.Name = model.Name
		stockItem.Update(context.Background(), r.db, boil.Infer())
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

package repository

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"openapi/internal/infra/sqlboiler"

	"openapi/internal/stockitem/model"
)

// Save saves the stock item model to the repository.
// It takes a StockItemModel as a parameter and returns an error.
func Save(db *sql.DB, model model.StockItemModel) error {

	exists, findErr := sqlboiler.FindStockItem(context.Background(), db, model.Id.String())
	stockItem := sqlboiler.StockItem{}
	if findErr != nil {
		stockItem.ID = model.Id.String()
		stockItem.Name = string(model.Name)
		dbExecErr := stockItem.Insert(context.Background(), db, boil.Infer())
		if dbExecErr != nil {
			return dbExecErr
		}
	} else {
		stockItem.ID = exists.ID
		stockItem.Name = model.Name
		stockItem.Update(context.Background(), db, boil.Infer())
	}

	return nil
}

func Get(db *sql.DB, id uuid.UUID) (model.StockItemModel, error) {
	exists, findErr := sqlboiler.FindStockItem(context.Background(), db, id.String())
	if findErr != nil {
		return model.StockItemModel{}, findErr
	}

	stockItem := model.Renew(
		uuid.MustParse(exists.ID),
		exists.Name,
	)
	return *stockItem, nil
}

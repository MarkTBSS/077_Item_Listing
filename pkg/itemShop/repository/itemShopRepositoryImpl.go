package repository

import (
	"github.com/MarkTBSS/077_Item_Listing/databases"
	"github.com/MarkTBSS/077_Item_Listing/entities"
	"github.com/labstack/echo/v4"
)

type itemRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db databases.Database, logger echo.Logger) ItemShopRepository {
	return &itemRepositoryImpl{
		db:     db,
		logger: logger,
	}
}
func (r *itemRepositoryImpl) Listing() ([]*entities.Item, error) {
	query := r.db.Connect()
	itemLists := make([]*entities.Item, 0)

	err := query.Find(&itemLists).Error
	if err != nil {
		r.logger.Error("Failed to find items", err.Error())
		return nil, err
	}
	return itemLists, nil
}

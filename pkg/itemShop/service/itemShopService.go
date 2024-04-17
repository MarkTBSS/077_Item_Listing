package service

import "github.com/MarkTBSS/077_Item_Listing/models"

type ItemShopService interface {
	Listing() ([]*models.Item, error)
}

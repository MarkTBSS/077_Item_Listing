package repository

import "github.com/MarkTBSS/077_Item_Listing/entities"

type ItemShopRepository interface {
	Listing() ([]*entities.Item, error)
}

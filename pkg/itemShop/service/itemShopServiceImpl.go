package service

import (
	"github.com/MarkTBSS/077_Item_Listing/models"
	"github.com/MarkTBSS/077_Item_Listing/pkg/itemShop/repository"
)

type itemShopServiceImpl struct {
	itemShopRepository repository.ItemShopRepository
}

func NewItemShopRepositoryImpl(itemShopRepository repository.ItemShopRepository) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository}
}

func (s *itemShopServiceImpl) Listing() ([]*models.Item, error) {
	itemEntityList, err := s.itemShopRepository.Listing()
	if err != nil {
		return nil, err
	}

	itemModelLists := make([]*models.Item, 0)
	for _, itemEntity := range itemEntityList {
		itemModelLists = append(itemModelLists, itemEntity.ChangeToItemModel())
	}

	return itemModelLists, nil
}

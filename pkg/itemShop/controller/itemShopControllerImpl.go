package controller

import (
	"net/http"

	"github.com/MarkTBSS/077_Item_Listing/pkg/itemShop/service"
	"github.com/labstack/echo/v4"
)

type itemShopControllerImpl struct {
	itemShopService service.ItemShopService
}

func NewItemShopControllerImpl(itemShopService service.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {
	itemModelLists, err := c.itemShopService.Listing()
	if err != nil {
		return pctx.String(http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusOK, itemModelLists)
}

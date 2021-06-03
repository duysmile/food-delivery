package cartbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/cart/cartmodel"
	"context"
)

type GetCartStore interface {
	ListItemInCartByUserId(ctx context.Context, userId int, paging *common.Paging, moreKeys ...string) ([]cartmodel.Cart, error)
}

type getCartBiz struct {
	store GetCartStore
}

func NewGetCartBiz(store GetCartStore) *getCartBiz {
	return &getCartBiz{store: store}
}

func (biz *getCartBiz) ListItemInCartByUserId(
	ctx context.Context,
	userId int,
	paging *common.Paging,
) ([]cartmodel.Cart, error) {
	items, err := biz.store.ListItemInCartByUserId(ctx, userId, paging, "Food")

	if err != nil {
		return nil, common.ErrCannotListEntity(cartmodel.EntityName, err)
	}

	return items, nil
}

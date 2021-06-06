package cartbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/cart/cartmodel"
	"context"
)

type UpdateCartStore interface {
	UpdateCart(ctx context.Context, userId int, foodId int, data *cartmodel.CartUpdate) error
}

type updateCartBiz struct {
	store UpdateCartStore
}

func NewUpdateCartBiz(store UpdateCartStore) *updateCartBiz {
	return &updateCartBiz{store: store}
}

func (biz *updateCartBiz) UpdateCart(ctx context.Context, userId int, foodId int, data *cartmodel.CartUpdate) error {
	if err := biz.store.UpdateCart(ctx, userId, foodId, data); err != nil {
		return common.ErrCannotUpdateEntity(cartmodel.EntityName, err)
	}

	return nil
}

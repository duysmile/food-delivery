package cartbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/cart/cartmodel"
	"context"
)

type DeleteCartStore interface {
	GetCart(ctx context.Context, userId int, foodId int) (*cartmodel.Cart, error)
	DeleteCart(ctx context.Context, cart *cartmodel.CartDelete) error
}

type deleteCartBiz struct {
	store DeleteCartStore
}

func NewDeleteCartBiz(store DeleteCartStore) *deleteCartBiz {
	return &deleteCartBiz{store: store}
}

func (biz *deleteCartBiz) DeleteCart(ctx context.Context, data *cartmodel.CartDelete) error {
	_, err := biz.store.GetCart(ctx, data.UserId, data.FoodId)

	if err != nil {
		return common.ErrCannotDeleteEntity(cartmodel.EntityName, err)
	}

	if err = biz.store.DeleteCart(ctx, data); err != nil {
		return common.ErrCannotDeleteEntity(cartmodel.EntityName, err)
	}

	return nil
}

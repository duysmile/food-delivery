package cartbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/cart/cartmodel"
	"context"
)

type CreateCartStore interface {
	GetCart(ctx context.Context, userId int, foodId int) (*cartmodel.Cart, error)
	CreateCart(ctx context.Context, data *cartmodel.CartCreate) error
	UpdateCart(ctx context.Context, userId int, foodId int, data *cartmodel.CartUpdate) error
}

type createCartBiz struct {
	store CreateCartStore
}

func NewCreateCartBiz(store CreateCartStore) *createCartBiz {
	return &createCartBiz{store: store}
}

func (biz *createCartBiz) CreateCart(ctx context.Context, data *cartmodel.CartCreate) error {
	userId := data.UserId
	foodId := data.FoodId

	cart, err := biz.store.GetCart(ctx, userId, foodId)
	if err != nil && err != common.RecordNotFound {
		return common.ErrCannotCreateEntity(cartmodel.EntityName, err)
	}

	if cart != nil {
		if err := biz.store.UpdateCart(ctx, userId, foodId, &cartmodel.CartUpdate{Quantity: data.Quantity}); err != nil {
			return common.ErrCannotCreateEntity(cartmodel.EntityName, err)
		}
	} else {
		if err := biz.store.CreateCart(ctx, data); err != nil {
			return common.ErrCannotCreateEntity(cartmodel.EntityName, err)
		}
	}

	return nil
}

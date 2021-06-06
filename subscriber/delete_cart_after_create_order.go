package subscriber

import (
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/cart/cartstorage"
	"200lab/food-delivery/pubsub"
	"context"
)

type HasCartDelete interface {
	GetUserId() int
	GetFoodIds() []int
}

func RunDeleteCartAfterCreateOrder(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Delete cart after create an order",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			store := cartstorage.NewSQLStore(appCtx.GetMainDBConnection())

			cartData := message.Data().(HasCartDelete)
			return store.DeleteCarts(ctx, cartData.GetUserId(), cartData.GetFoodIds())
		},
	}
}

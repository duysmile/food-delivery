package subscriber

import (
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/ordertracking/ordertrackingmodel"
	"200lab/food-delivery/modules/ordertracking/ordertrackingstorage"
	"200lab/food-delivery/pubsub"
	"context"
)

type HasOrderTrackingData interface {
	GetOrderId() int
	GetState() ordertrackingmodel.OrderState
}

func RunCreateOrderTrackingAfterCreateOrder(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Create order tracking after create order",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			store := ordertrackingstorage.NewSQLStore(appCtx.GetMainDBConnection())

			data := message.Data().(HasOrderTrackingData)
			return store.CreateOrderTracking(ctx, &ordertrackingmodel.OrderTrackingCreate{
				OrderId: data.GetOrderId(),
				State:   data.GetState(),
			})
		},
	}
}

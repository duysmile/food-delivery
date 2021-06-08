package ordertrackingbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/order/ordermodel"
	"200lab/food-delivery/modules/ordertracking/ordertrackingmodel"
	"context"
)

type CancelOrderStore interface {
	UpdateOrderTracking(ctx context.Context, id int, data *ordertrackingmodel.OrderTrackingUpdate) error
	GetOrderTrackingByCondition(ctx context.Context, id int, userId int, condition map[string]interface{}) (*ordertrackingmodel.OrderTracking, error)
}

type GetOrderStore interface {
	GetOrder(ctx context.Context, id int, condition map[string]interface{}) (*ordermodel.Order, error)
}

type cancelOrderBiz struct {
	store      CancelOrderStore
	orderStore GetOrderStore
}

func NewCancelOrderBiz(store CancelOrderStore, orderStore GetOrderStore) *cancelOrderBiz {
	return &cancelOrderBiz{store: store, orderStore: orderStore}
}

func (biz *cancelOrderBiz) CancelOrderStore(ctx context.Context, userId int, orderId int) error {
	orderTracking, err := biz.store.GetOrderTrackingByCondition(ctx, orderId, userId, map[string]interface{}{
		"state": ordertrackingmodel.Preparing,
	})

	if err != nil {
		return common.ErrCannotUpdateEntity(ordertrackingmodel.EntityName, err)
	}

	dataUpdate := ordertrackingmodel.OrderTrackingUpdate{
		State: ordertrackingmodel.Cancel,
	}
	if err = biz.store.UpdateOrderTracking(ctx, orderTracking.Id, &dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(ordertrackingmodel.EntityName, err)
	}

	return nil
}

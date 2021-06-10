package ordertrackingbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/order/ordermodel"
	"200lab/food-delivery/modules/ordertracking/ordertrackingmodel"
	"context"
)

type CreateOrderTrackingStore interface {
	CreateOrderTracking(ctx context.Context, data *ordertrackingmodel.OrderTrackingCreate) error
}

type GetOrderStore interface {
	GetOrder(ctx context.Context, id int, condition map[string]interface{}) (*ordermodel.Order, error)
}

type createOrderTrackingBiz struct {
	store      CreateOrderTrackingStore
	orderStore GetOrderStore
}

func NewCreateOrderBiz(store CreateOrderTrackingStore, orderStore GetOrderStore) *createOrderTrackingBiz {
	return &createOrderTrackingBiz{store: store, orderStore: orderStore}
}

func (biz *createOrderTrackingBiz) CreateOrderTracking(
	ctx context.Context,
	userId int,
	orderTracking *ordertrackingmodel.OrderTrackingCreate,
) error {
	_, err := biz.orderStore.GetOrder(ctx, orderTracking.OrderId, map[string]interface{}{
		"user_id": userId,
	})

	if err != nil {
		return common.ErrCannotCreateEntity(ordertrackingmodel.EntityName, err)
	}

	// TODO check if cancel -> order must not be in preparing

	if err = biz.store.CreateOrderTracking(ctx, orderTracking); err != nil {
		return common.ErrCannotCreateEntity(ordertrackingmodel.EntityName, err)
	}

	return nil
}

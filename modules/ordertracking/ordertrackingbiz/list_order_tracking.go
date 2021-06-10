package ordertrackingbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/ordertracking/ordertrackingmodel"
	"context"
)

type ListOrderTrackingStore interface {
	ListOrderTrackingByCondition(
		ctx context.Context,
		condition map[string]interface{},
		paging *common.Paging,
		moreKeys ...string,
	) ([]ordertrackingmodel.OrderTracking, error)
}

type listOrderTrackingBiz struct {
	store      ListOrderTrackingStore
	orderStore GetOrderStore
}

func NewListOrderTrackingBiz(store ListOrderTrackingStore, orderStore GetOrderStore) *listOrderTrackingBiz {
	return &listOrderTrackingBiz{store: store, orderStore: orderStore}
}

func (biz *listOrderTrackingBiz) ListOrderTracking(ctx context.Context, userId int, orderId int, paging *common.Paging) ([]ordertrackingmodel.OrderTracking, error) {
	if _, err := biz.orderStore.GetOrder(ctx, orderId, map[string]interface{}{"user_id": userId}); err != nil {
		return nil, common.ErrCannotListEntity(ordertrackingmodel.EntityName, err)
	}

	orderTrackings, err := biz.store.ListOrderTrackingByCondition(ctx, map[string]interface{}{
		"order_id": orderId,
	}, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(ordertrackingmodel.EntityName, err)
	}

	return orderTrackings, nil
}

package ordertrackingbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/ordertracking/ordertrackingmodel"
	"context"
)

type GetOrderTrackingStore interface {
	GetOrderTrackingByCondition(ctx context.Context, id int, userId int, notCondition map[string]interface{}) (*ordertrackingmodel.OrderTracking, error)
}

type getOrderTrackingBiz struct {
	store GetOrderTrackingStore
}

func NewGetOrderTrackingBiz(store GetOrderTrackingStore) *getOrderTrackingBiz {
	return &getOrderTrackingBiz{store: store}
}

func (biz *getOrderTrackingBiz) GetOrderTracking(ctx context.Context, userId int, id int) (*ordertrackingmodel.OrderTracking, error) {
	orderTracking, err := biz.store.GetOrderTrackingByCondition(ctx, id, userId, nil)

	if err != nil {
		return nil, common.ErrCannotGetEntity(ordertrackingmodel.EntityName, err)
	}

	return orderTracking, nil
}

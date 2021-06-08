package ordertrackingbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/ordertracking/ordertrackingmodel"
	"context"
	"time"
)

type CancelOrderStore interface {
	UpdateOrderTracking(ctx context.Context, id int, data *ordertrackingmodel.OrderTrackingUpdate) error
	GetOrderTrackingByCondition(ctx context.Context, id int, notCondition map[string]interface{}) (*ordertrackingmodel.OrderTracking, error)
}

type cancelOrderBiz struct {
	store CancelOrderStore
}

func NewCancelOrderBiz(store CancelOrderStore) *cancelOrderBiz {
	return &cancelOrderBiz{store: store}
}

func (biz *cancelOrderBiz) CancelOrderStore(ctx context.Context, userId int, id int) error {
	_, err := biz.store.GetOrderTrackingByCondition(ctx, id, map[string]interface{}{
		"state": []ordertrackingmodel.OrderState{
			ordertrackingmodel.Deliveried,
			ordertrackingmodel.OnTheWay,
		},
		"created_at > ?": time.Now().Add(time.Minute * time.Duration(3)).Format("2006-01-02 15:04:05"),
	})

	if err != nil {
		return common.ErrCannotUpdateEntity(ordertrackingmodel.EntityName, err)
	}

	dataUpdate := ordertrackingmodel.OrderTrackingUpdate{
		State: ordertrackingmodel.Cancel,
	}
	if err = biz.store.UpdateOrderTracking(ctx, id, &dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(ordertrackingmodel.EntityName, err)
	}

	return nil
}

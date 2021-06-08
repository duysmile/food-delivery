package orderbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/order/ordermodel"
	"context"
)

type ListOrderStore interface {
	ListOrdersByCondition(
		ctx context.Context,
		condition map[string]interface{},
		paging *common.Paging,
		moreKeys ...string,
	) ([]ordermodel.Order, error)
}

type listOrderBiz struct {
	store ListOrderStore
}

func NewListOrderBiz(store ListOrderStore) *listOrderBiz {
	return &listOrderBiz{store: store}
}

func (biz *listOrderBiz) ListOrder(ctx context.Context, userId int, paging *common.Paging) ([]ordermodel.Order, error) {
	orders, err := biz.store.ListOrdersByCondition(ctx, map[string]interface{}{
		"user_id": userId,
	}, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(ordermodel.EntityName, err)
	}

	return orders, nil
}

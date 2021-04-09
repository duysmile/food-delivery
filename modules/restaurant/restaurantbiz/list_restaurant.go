package restaurantbiz

import (
	"200lab/food-delivery/common"
	restarantmodel "200lab/food-delivery/modules/restaurant/restaurantmodel"
	"context"
)

type ListUserStore interface {
	ListByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *restarantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restarantmodel.Restaurant, error)
}

type listUserBiz struct {
	store ListUserStore
}

func NewListUserBiz(store ListUserStore) *listUserBiz {
	return &listUserBiz{store: store}
}

func (biz *listUserBiz) ListUserBiz(ctx context.Context,
	conditions map[string]interface{},
	filter *restarantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restarantmodel.Restaurant, error) {
	var result []restarantmodel.Restaurant

	result, err := biz.store.ListByCondition(ctx, nil, filter, paging)
	return result, err
}

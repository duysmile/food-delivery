package restaurantbiz

import (
	"200lab/food-delivery/common"
	restaurantmodel "200lab/food-delivery/modules/restaurant/restaurantmodel"
	"context"
)

type ListUserStore interface {
	ListByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type listUserBiz struct {
	store ListUserStore
}

func NewListUserBiz(store ListUserStore) *listUserBiz {
	return &listUserBiz{store: store}
}

func (biz *listUserBiz) ListUserBiz(ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant

	result, err := biz.store.ListByCondition(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	return result, nil
}

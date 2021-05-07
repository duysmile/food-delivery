package restaurantbiz

import (
	"200lab/food-delivery/common"
	restaurantmodel "200lab/food-delivery/modules/restaurant/restaurantmodel"
	"context"
	"log"
)

type ListUserStore interface {
	ListByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type LikeStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listUserBiz struct {
	store     ListUserStore
	likeStore LikeStore
}

func NewListUserBiz(store ListUserStore, likeStore LikeStore) *listUserBiz {
	return &listUserBiz{store: store, likeStore: likeStore}
}

func (biz *listUserBiz) ListUserBiz(ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant

	result, err := biz.store.ListByCondition(ctx, nil, filter, paging, "Owner")
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	mapResLike, err := biz.likeStore.GetRestaurantLikes(ctx, ids)
	if err != nil {
		log.Println("cannot get restaurant likes:", err)
	}

	if v := mapResLike; v != nil {
		for i, item := range result {
			result[i].LikeCount = mapResLike[item.Id]
		}
	}

	return result, nil
}

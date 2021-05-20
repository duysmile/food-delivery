package foodbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/food/foodmodel"
	"context"
)

type ListFoodStore interface {
	ListFoodByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *foodmodel.Filter,
		paging *common.Paging,
		moreInfo ...string,
	) ([]foodmodel.Food, error)
}

type listFoodBiz struct {
	store ListFoodStore
}

func NewListFoodBiz(store ListFoodStore) *listFoodBiz {
	return &listFoodBiz{store: store}
}

func (biz *listFoodBiz) ListFoodByRestaurantId(
	ctx context.Context,
	restaurantId int,
	paging *common.Paging,
) ([]foodmodel.Food, error) {
	store := biz.store

	data, err := store.ListFoodByCondition(ctx, map[string]interface{}{"restaurant_id": restaurantId}, nil, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(foodmodel.EntityName, err)
	}

	return data, nil
}

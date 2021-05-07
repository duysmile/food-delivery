package foodbiz

import (
	"200lab/food-delivery/modules/food/foodmodel"
	"200lab/food-delivery/modules/restaurant/restaurantmodel"
	"context"
)

type CreateDataStore interface {
	CreateFood(ctx context.Context, data *foodmodel.FoodCreate) error
}

type GetRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type createFoodBiz struct {
	store           CreateDataStore
	restaurantStore GetRestaurantStore
}

func NewCreateFoodBiz(store CreateDataStore, restaurantStore GetRestaurantStore) *createFoodBiz {
	return &createFoodBiz{store: store, restaurantStore: restaurantStore}
}

func (biz *createFoodBiz) CreateFood(ctx context.Context, userId int, data *foodmodel.FoodCreate) error {
	_, err := biz.restaurantStore.FindDataByCondition(ctx, map[string]interface{}{
		"id":       data.RestaurantId,
		"owner_id": userId,
	})
	if err != nil {
		return err
	}

	if err := biz.store.CreateFood(ctx, data); err != nil {
		return err
	}

	return nil
}

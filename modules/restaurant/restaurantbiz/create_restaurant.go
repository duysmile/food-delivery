package restaurantbiz

import (
	restarantmodel "200lab/food-delivery/modules/restaurant/restaurantmodel"
	"context"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restarantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restarantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	err := biz.store.Create(ctx, data)
	return err
}

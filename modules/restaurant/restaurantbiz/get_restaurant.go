package restaurantbiz

import (
	restarantmodel "200lab/food-delivery/modules/restaurant/restaurantmodel"
	"context"
)

type GetRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restarantmodel.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantStore(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant(ctx context.Context, id int) (*restarantmodel.Restaurant, error) {
	store := biz.store
	var data *restarantmodel.Restaurant

	data, err := store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	return data, err
}

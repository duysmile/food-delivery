package restaurantbiz

import (
	"200lab/food-delivery/common"
	restaurantmodel "200lab/food-delivery/modules/restaurant/restaurantmodel"
	"context"
	"errors"
)

type UpdateRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	UpdateRestaurant(
		ctx context.Context,
		id int,
		data *restaurantmodel.RestaurantUpdate,
	) error
}

type updateRestaurantBiz struct {
	store      UpdateRestaurantStore
	imageStore FindImageStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore, imageStore FindImageStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store, imageStore: imageStore}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(
	ctx context.Context,
	id int,
	userId int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	store := biz.store

	oldData, err := store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityNotFound(restaurantmodel.EntityName, nil)
	}

	if userId != oldData.OwnerId {
		panic(common.ErrNoPermission(errors.New("userId is not match ownerId")))
	}

	if err := data.Logo.Validate(ctx, biz.imageStore); err != nil {
		panic(err)
	}

	if err := data.Cover.Validate(ctx, biz.imageStore); err != nil {
		panic(err)
	}

	if err := store.UpdateRestaurant(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}

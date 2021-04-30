package restaurantbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/restaurant/restaurantmodel"
	"context"
	"errors"
)

type DeleteRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	SoftDelete(ctx context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(
	ctx context.Context,
	id int,
	userId int,
) error {
	store := biz.store

	oldData, err := store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	if userId != oldData.OwnerId {
		panic(common.ErrNoPermission(errors.New("userId is not match ownerId")))
	}

	if err := store.SoftDelete(ctx, id); err != nil {
		return err
	}

	return nil
}

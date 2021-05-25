package foodbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/food/foodmodel"
	"context"
	"errors"
)

type DeleteFoodStore interface {
	DeleteFood(ctx context.Context, id int) error
	FindFoodById(ctx context.Context, id int, moreInfo ...string) (*foodmodel.Food, error)
}

type deleteFoodBiz struct {
	store DeleteFoodStore
}

func NewDeleteFoodBiz(store DeleteFoodStore) *deleteFoodBiz {
	return &deleteFoodBiz{store: store}
}

func (biz *deleteFoodBiz) DeleteFood(ctx context.Context, userId int, id int) error {
	food, err := biz.store.FindFoodById(ctx, id, "Restaurant")
	if err != nil {
		return common.ErrCannotDeleteEntity(foodmodel.EntityName, err)
	}

	ownerId := food.Restaurant.OwnerId
	if ownerId != userId {
		return common.ErrNoPermission(errors.New("user not own this food"))
	}

	if err = biz.store.DeleteFood(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(foodmodel.EntityName, err)
	}

	return nil
}

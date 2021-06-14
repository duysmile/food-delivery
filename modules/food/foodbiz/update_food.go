package foodbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/food/foodmodel"
	"context"
	"errors"
)

type UpdateFoodStore interface {
	UpdateFoodById(ctx context.Context, id int, data *foodmodel.FoodUpdate) error
	FindFoodById(ctx context.Context, id int, moreInfo ...string) (*foodmodel.Food, error)
}

type updateFoodBiz struct {
	store UpdateFoodStore
}

func NewUpdateFoodBiz(store UpdateFoodStore) *updateFoodBiz {
	return &updateFoodBiz{store: store}
}

func (biz *updateFoodBiz) UpdateFood(ctx context.Context, userId int, id int, data *foodmodel.FoodUpdate) error {
	food, err := biz.store.FindFoodById(ctx, id, "Restaurant")

	if err != nil {
		return common.ErrCannotUpdateEntity(foodmodel.EntityName, err)
	}

	if food.Restaurant == nil {
		return common.ErrEntityNotFound(foodmodel.EntityName, errors.New("food not belong any restaurant"))
	}

	ownerId := food.Restaurant.OwnerId
	if ownerId != userId {
		return common.ErrNoPermission(errors.New("user not own food"))
	}

	if err = biz.store.UpdateFoodById(ctx, id, data); err != nil {
		return err
	}

	return nil
}

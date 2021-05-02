package restaurantlikebiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/restaurantlike/restaurantlikemodel"
	"context"
)

type CreateRestaurantLikeStore interface {
	CreateRestaurantLike(ctx context.Context, data *restaurantlikemodel.LikeCreate) error
	GetRestaurantLike(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*restaurantlikemodel.Like, error)
}

type createRestaurantLikeBiz struct {
	createStore CreateRestaurantLikeStore
}

func NewCreateRestaurantLikeBiz(
	createStore CreateRestaurantLikeStore,
) *createRestaurantLikeBiz {
	return &createRestaurantLikeBiz{
		createStore: createStore,
	}
}

func (biz *createRestaurantLikeBiz) CreateLike(ctx context.Context, data *restaurantlikemodel.LikeCreate) error {
	oldLike, err := biz.createStore.GetRestaurantLike(ctx, map[string]interface{}{
		"restaurant_id": data.RestaurantId,
		"user_id":       data.UserId,
	})

	if err != nil && err != common.RecordNotFound {
		return err
	}

	if oldLike != nil {
		return nil
	}

	if err := biz.createStore.CreateRestaurantLike(ctx, data); err != nil {
		return err
	}

	return nil
}

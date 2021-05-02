package restaurantlikebiz

import (
	"200lab/food-delivery/modules/restaurantlike/restaurantlikemodel"
	"context"
)

type DeleteRestaurantLikeStore interface {
	DeleteRestaurantLike(ctx context.Context, likeDelete *restaurantlikemodel.LikeDelete) error
	GetRestaurantLike(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*restaurantlikemodel.Like, error)
}

type deleteRestaurantLikeBiz struct {
	deleteStore DeleteRestaurantLikeStore
}

func NewDeleteRestaurantLikeBiz(
	deleteStore DeleteRestaurantLikeStore,
) *deleteRestaurantLikeBiz {
	return &deleteRestaurantLikeBiz{
		deleteStore: deleteStore,
	}
}

func (biz *deleteRestaurantLikeBiz) DeleteLike(ctx context.Context, data *restaurantlikemodel.LikeDelete) error {
	oldLike, _ := biz.deleteStore.GetRestaurantLike(ctx, map[string]interface{}{
		"restaurant_id": data.RestaurantId,
		"user_id":       data.UserId,
	})

	if oldLike == nil {
		return nil
	}

	if err := biz.deleteStore.DeleteRestaurantLike(ctx, data); err != nil {
		return err
	}

	return nil
}

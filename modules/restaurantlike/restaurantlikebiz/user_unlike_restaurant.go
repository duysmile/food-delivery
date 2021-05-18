package restaurantlikebiz

import (
	"200lab/food-delivery/component/asyncjob"
	"200lab/food-delivery/modules/restaurantlike/restaurantlikemodel"
	"context"
)

type DeleteRestaurantLikeStore interface {
	DeleteRestaurantLike(ctx context.Context, likeDelete *restaurantlikemodel.LikeDelete) error
	GetRestaurantLike(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*restaurantlikemodel.Like, error)
}

type DecreaseLikedCountStore interface {
	DecreaseLikedCount(ctx context.Context, id int) error
}

type deleteRestaurantLikeBiz struct {
	deleteStore             DeleteRestaurantLikeStore
	decreaseLikedCountStore DecreaseLikedCountStore
}

func NewDeleteRestaurantLikeBiz(
	deleteStore DeleteRestaurantLikeStore,
	decreaseLikedCountStore DecreaseLikedCountStore,
) *deleteRestaurantLikeBiz {
	return &deleteRestaurantLikeBiz{
		deleteStore:             deleteStore,
		decreaseLikedCountStore: decreaseLikedCountStore,
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

	// side effect
	job := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.decreaseLikedCountStore.DecreaseLikedCount(ctx, data.RestaurantId)
	})

	asyncjob.NewGroup(true, job).Run(ctx)

	return nil
}

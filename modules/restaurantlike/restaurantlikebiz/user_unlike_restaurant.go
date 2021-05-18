package restaurantlikebiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/restaurantlike/restaurantlikemodel"
	"200lab/food-delivery/pubsub"
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
	deleteStore DeleteRestaurantLikeStore
	// decreaseLikedCountStore DecreaseLikedCountStore
	pubsub pubsub.Pubsub
}

func NewDeleteRestaurantLikeBiz(
	deleteStore DeleteRestaurantLikeStore,
	// decreaseLikedCountStore DecreaseLikedCountStore,
	pb pubsub.Pubsub,
) *deleteRestaurantLikeBiz {
	return &deleteRestaurantLikeBiz{
		deleteStore: deleteStore,
		// decreaseLikedCountStore: decreaseLikedCountStore,
		pubsub: pb,
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
	// use async job
	// job := asyncjob.NewJob(func(ctx context.Context) error {
	// 	return biz.decreaseLikedCountStore.DecreaseLikedCount(ctx, data.RestaurantId)
	// })

	// asyncjob.NewGroup(true, job).Run(ctx)

	// use pubsub
	biz.pubsub.Publish(ctx, common.TopicUserUnLikeRestaurant, pubsub.NewMessage(data))

	return nil
}

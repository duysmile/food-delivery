package restaurantlikebiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/restaurantlike/restaurantlikemodel"
	"200lab/food-delivery/pubsub"
	"context"
)

type CreateRestaurantLikeStore interface {
	CreateRestaurantLike(ctx context.Context, data *restaurantlikemodel.LikeCreate) error
	GetRestaurantLike(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*restaurantlikemodel.Like, error)
}

type IncreaseLikedCountStore interface {
	IncreaseLikedCount(ctx context.Context, id int) error
}

type createRestaurantLikeBiz struct {
	createStore CreateRestaurantLikeStore
	// increaseLikedCountStore IncreaseLikedCountStore
	pubsub pubsub.Pubsub
}

func NewCreateRestaurantLikeBiz(
	createStore CreateRestaurantLikeStore,
	// increaseLikedCountStore IncreaseLikedCountStore,
	pb pubsub.Pubsub,
) *createRestaurantLikeBiz {
	return &createRestaurantLikeBiz{
		createStore: createStore,
		// increaseLikedCountStore: increaseLikedCountStore,
		pubsub: pb,
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

	// side effect

	// use async job
	// job := asyncjob.NewJob(func(ctx context.Context) error {
	// 	return biz.increaseLikedCountStore.IncreaseLikedCount(ctx, data.RestaurantId)
	// })

	// group := asyncjob.NewGroup(true, job)
	// group.Run(ctx)

	// use pub/sub
	biz.pubsub.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data))

	return nil
}

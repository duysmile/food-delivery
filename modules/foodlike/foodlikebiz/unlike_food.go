package foodlikebiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/foodlike/foodlikemodel"
	"200lab/food-delivery/pubsub"
	"context"
)

type UnlikeFoodStore interface {
	UnlikeFood(ctx context.Context, data *foodlikemodel.FoodLikeCreate) error
	GetLikeFood(ctx context.Context, userId int, foodId int, moreKeys ...string) (*foodlikemodel.FoodLike, error)
}

type unlikeFoodBiz struct {
	store  UnlikeFoodStore
	pubsub pubsub.Pubsub
}

func NewUnlikeFoodBiz(store UnlikeFoodStore, pb pubsub.Pubsub) *unlikeFoodBiz {
	return &unlikeFoodBiz{
		store:  store,
		pubsub: pb,
	}
}

func (biz *unlikeFoodBiz) UnlikeFood(ctx context.Context, data *foodlikemodel.FoodLikeCreate) error {
	oldLike, err := biz.store.GetLikeFood(ctx, data.UserId, data.FoodId)

	if err != nil && err != common.RecordNotFound {
		return err
	}

	if oldLike == nil {
		return nil
	}

	if err := biz.store.UnlikeFood(ctx, data); err != nil {
		return common.ErrCannotDeleteEntity(foodlikemodel.EntityName, err)
	}

	// publish event to decrease like count of food
	biz.pubsub.Publish(ctx, common.TopicUserUnLikeFood, pubsub.NewMessage(data))

	return nil
}

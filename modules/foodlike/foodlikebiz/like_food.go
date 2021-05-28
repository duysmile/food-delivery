package foodlikebiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/foodlike/foodlikemodel"
	"200lab/food-delivery/pubsub"
	"context"
)

type LikeFoodStore interface {
	LikeFood(ctx context.Context, data *foodlikemodel.FoodLikeCreate) error
	GetLikeFood(ctx context.Context, userId int, foodId int, moreKeys ...string) (*foodlikemodel.FoodLike, error)
}

type likeFoodBiz struct {
	store  LikeFoodStore
	pubsub pubsub.Pubsub
}

func NewLikeFoodBiz(
	store LikeFoodStore,
	pb pubsub.Pubsub,
) *likeFoodBiz {
	return &likeFoodBiz{
		store:  store,
		pubsub: pb,
	}
}

func (biz *likeFoodBiz) LikeFood(ctx context.Context, data *foodlikemodel.FoodLikeCreate) error {
	oldLike, err := biz.store.GetLikeFood(ctx, data.UserId, data.FoodId)

	if err != nil && err != common.RecordNotFound {
		return err
	}

	if oldLike != nil {
		return nil
	}

	if err := biz.store.LikeFood(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(foodlikemodel.EntityName, err)
	}

	// publish event increase food like in food model
	biz.pubsub.Publish(ctx, common.TopicUserLikeFood, pubsub.NewMessage(data))

	return nil
}

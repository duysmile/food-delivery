package subscriber

import (
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/food/foodstorage"
	"200lab/food-delivery/pubsub"
	"context"
)

type HasFoodId interface {
	GetFoodId() int
}

func RunIncreaseLikeCountAfterUserLikeFood(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after an user like food",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			store := foodstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasFoodId)

			return store.IncreaseLikedCount(ctx, likeData.GetFoodId())
		},
	}
}

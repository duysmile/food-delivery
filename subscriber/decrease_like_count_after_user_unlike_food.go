package subscriber

import (
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/food/foodstorage"
	"200lab/food-delivery/pubsub"
	"context"
)

func RunDecreaseLikeCountAfterUserUnLikeFood(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user unlike food",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			store := foodstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasFoodId)
			return store.DecreaseLikedCount(ctx, likeData.GetFoodId())
		},
	}
}

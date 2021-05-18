package subscriber

import (
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/restaurant/restaurantstorage"
	"200lab/food-delivery/pubsub"
	"context"
)

func RunDecreaseLikeCountAfterUserUnLikeRestaurant(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user like restaurant",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.IncreaseLikedCount(ctx, likeData.GetRestaurantId())
		},
	}
}

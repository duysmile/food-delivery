package restaurantbiz

import (
	"200lab/food-delivery/common"
	restaurantmodel "200lab/food-delivery/modules/restaurant/restaurantmodel"
	"context"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type FindImageStore interface {
	FindImageByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*common.Image, error)
	ListImages(
		ctx context.Context,
		ids []int,
		moreKeys ...string,
	) ([]common.Image, error)
}

type createRestaurantBiz struct {
	store      CreateRestaurantStore
	imageStore FindImageStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore, imageStore FindImageStore) *createRestaurantBiz {
	return &createRestaurantBiz{
		store:      store,
		imageStore: imageStore,
	}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := data.Logo.Validate(ctx, biz.imageStore); err != nil {
		return err
	}

	if err := data.Cover.Validate(ctx, biz.imageStore); err != nil {
		return err
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}

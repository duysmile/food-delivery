package restaurantlikestorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/restaurantlike/restaurantlikemodel"
	"context"
)

func (s *sqlStore) CreateRestaurantLike(ctx context.Context, data *restaurantlikemodel.LikeCreate) error {
	db := s.db.Table(restaurantlikemodel.LikeCreate{}.TableName())

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

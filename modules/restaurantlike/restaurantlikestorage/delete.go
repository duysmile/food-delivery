package restaurantlikestorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/restaurantlike/restaurantlikemodel"
	"context"
)

func (s *sqlStore) DeleteRestaurantLike(ctx context.Context, likeDelete *restaurantlikemodel.LikeDelete) error {
	db := s.db.Table(restaurantlikemodel.LikeDelete{}.TableName())

	if err := db.Delete(likeDelete).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

package restaurantlikestorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/restaurantlike/restaurantlikemodel"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) GetRestaurantLike(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*restaurantlikemodel.Like, error) {
	db := s.db.Table(restaurantlikemodel.Like{}.TableName())

	var like restaurantlikemodel.Like
	if err := db.Where(conditions).First(&like).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &like, nil
}

package restaurantstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/restaurant/restaurantmodel"
	"context"
)

func (s *sqlStore) SoftDelete(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

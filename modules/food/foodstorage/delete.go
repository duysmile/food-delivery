package foodstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/food/foodmodel"
	"context"
)

func (s *sqlStore) DeleteFood(ctx context.Context, id int) error {
	db := s.db.Table(foodmodel.Food{}.TableName())

	if err := db.Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

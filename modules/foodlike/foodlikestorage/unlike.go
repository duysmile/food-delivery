package foodlikestorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/foodlike/foodlikemodel"
	"context"
)

func (s *sqlStore) UnlikeFood(ctx context.Context, data *foodlikemodel.FoodLikeCreate) error {
	db := s.db.Table(foodlikemodel.FoodLikeCreate{}.TableName())

	if err := db.Delete(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

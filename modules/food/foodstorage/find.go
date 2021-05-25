package foodstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/food/foodmodel"
	"context"
)

func (s *sqlStore) FindFoodById(ctx context.Context, id int, moreInfo ...string) (*foodmodel.Food, error) {
	db := s.db.Table(foodmodel.Food{}.TableName())

	for i := range moreInfo {
		db.Preload(moreInfo[i])
	}

	var food foodmodel.Food
	if err := db.Where("id = ?", id).Find(&food).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &food, nil
}

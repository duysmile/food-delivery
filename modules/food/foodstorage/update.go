package foodstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/food/foodmodel"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) UpdateFoodById(ctx context.Context, id int, data *foodmodel.FoodUpdate) error {
	db := s.db.Table(foodmodel.Food{}.TableName()).Where("status in (?)", 1)

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseLikedCount(ctx context.Context, id int) error {
	db := s.db.Table(foodmodel.Food{}.TableName()).Where("status in (?)", 1)

	if err := db.Where("id = ?", id).Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreaseLikedCount(ctx context.Context, id int) error {
	db := s.db.Table(foodmodel.Food{}.TableName()).Where("status in (?)", 1)

	if err := db.Where("id = ?", id).Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

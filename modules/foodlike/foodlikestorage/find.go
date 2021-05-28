package foodlikestorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/foodlike/foodlikemodel"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) GetLikeFood(ctx context.Context, userId int, foodId int, moreKeys ...string) (*foodlikemodel.FoodLike, error) {
	db := s.db.Table(foodlikemodel.FoodLike{}.TableName())

	var foodlike foodlikemodel.FoodLike

	for i := range moreKeys {
		db.Preload(moreKeys[i])
	}

	if err := db.First(&foodlike).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &foodlike, nil
}

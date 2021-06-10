package categorystorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/category/categorymodel"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) GetCategoryByCondition(ctx context.Context, condition map[string]interface{}) (*categorymodel.Category, error) {
	db := s.db.Table(categorymodel.Category{}.TableName()).Where("status in (?)", 1)

	var data categorymodel.Category
	if err := db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}

package categorystorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/category/categorymodel"
	"context"
)

func (s *sqlStore) CreateCategory(ctx context.Context, data *categorymodel.CategoryCreate) error {
	db := s.db.Table(categorymodel.CategoryCreate{}.TableName())

	if err := db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

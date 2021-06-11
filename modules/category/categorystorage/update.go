package categorystorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/category/categorymodel"
	"context"
)

func (s *sqlStore) UpdateCategory(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error {
	db := s.db.Table(categorymodel.CategoryUpdate{}.TableName())

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

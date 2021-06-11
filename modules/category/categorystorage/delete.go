package categorystorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/category/categorymodel"
	"context"
)

func (s *sqlStore) DeleteCategory(ctx context.Context, id int) error {
	db := s.db.Table(categorymodel.Category{}.TableName()).Where("status in (?)", 1)

	if err := db.Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

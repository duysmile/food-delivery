package uploadstorage

import (
	"200lab/food-delivery/common"
	"context"
)

func (store *sqlStore) DeleteImage(ctx context.Context, ids []int) error {
	db := store.db

	if err := db.Table(common.Image{}.TableName()).
		Where("id in (?)", ids).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

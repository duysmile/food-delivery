package uploadstorage

import (
	"200lab/food-delivery/common"
	"context"
)

func (store *sqlStore) CreateImage(ctx context.Context, upload *common.Image) error {
	db := store.db

	if err := db.Create(upload).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

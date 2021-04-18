package uploadstorage

import (
	"200lab/food-delivery/common"
	"context"
)

func (store *sqlStore) ListImages(
	ctx context.Context,
	ids []int,
	moreKeys ...string,
) ([]common.Image, error) {
	db := store.db

	var images []common.Image
	if err := db.Table(common.Image{}.TableName()).
		Where("id in (?)", ids).
		Find(&images).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return images, nil
}

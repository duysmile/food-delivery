package uploadstorage

import (
	"200lab/food-delivery/common"
	"context"
)

func (s *sqlStore) FindImageByCondition(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*common.Image, error) {
	db := s.db.Table(common.Image{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var image common.Image

	if err := db.Where(conditions).First(&image).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &image, nil
}

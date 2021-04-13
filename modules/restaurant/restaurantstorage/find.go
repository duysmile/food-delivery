package restaurantstorage

import (
	"200lab/food-delivery/common"
	restaurantmodel "200lab/food-delivery/modules/restaurant/restaurantmodel"
	"context"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var data restaurantmodel.Restaurant
	if err := db.Where(conditions).First(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &data, nil
}

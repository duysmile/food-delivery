package restaurantstorage

import (
	restarantmodel "200lab/food-delivery/modules/restaurant/restaurantmodel"
	"context"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*restarantmodel.Restaurant, error) {
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var data restarantmodel.Restaurant
	if err := db.Where(conditions).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

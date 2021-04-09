package restaurantstorage

import (
	"200lab/food-delivery/common"
	restarantmodel "200lab/food-delivery/modules/restaurant/restaurantmodel"
	"context"
)

func (s *sqlStore) ListByCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *restarantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restarantmodel.Restaurant, error) {
	db := s.db

	for key := range moreKeys {
		db = db.Preload(moreKeys[key])
	}

	db = db.Table(restarantmodel.Restaurant{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.CityId > 0 {
			db = db.Where("city_id=?", v.CityId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	var result []restarantmodel.Restaurant
	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).
		Error; err != nil {
		return nil, err
	}

	return result, nil
}

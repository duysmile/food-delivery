package restaurantstorage

import (
	"200lab/food-delivery/common"
	restaurantmodel "200lab/food-delivery/modules/restaurant/restaurantmodel"
	"context"
)

func (s *sqlStore) UpdateRestaurant(
	ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

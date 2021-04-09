package restaurantstorage

import (
	restarantmodel "200lab/food-delivery/modules/restaurant/restaurantmodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *restarantmodel.RestaurantCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

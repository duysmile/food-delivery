package cartstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/cart/cartmodel"
	"context"
)

func (s *sqlStore) UpdateCart(ctx context.Context, userId int, foodId int, data *cartmodel.CartUpdate) error {
	db := s.db.Table(cartmodel.CartUpdate{}.TableName()).Where("status in (?)", 1)

	if err := db.Where(map[string]interface{}{
		"user_id": userId,
		"food_id": foodId,
	}).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

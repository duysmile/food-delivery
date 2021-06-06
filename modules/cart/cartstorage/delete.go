package cartstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/cart/cartmodel"
	"context"
)

func (s *sqlStore) DeleteCart(ctx context.Context, cart *cartmodel.CartDelete) error {
	db := s.db.Table(cartmodel.CartDelete{}.TableName()).Where("status in (?)", 1)

	if err := db.Delete(cart).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DeleteCarts(ctx context.Context, userId int, foodIds []int) error {
	db := s.db.Table(cartmodel.CartDelete{}.TableName()).Where("status in (?)", 1)

	if err := db.Where("user_id = ?", userId).
		Where("food_id in (?)", foodIds).
		Delete(&cartmodel.CartDelete{}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

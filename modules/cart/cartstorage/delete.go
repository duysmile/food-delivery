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

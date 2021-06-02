package cartstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/cart/cartmodel"
	"context"
)

func (s *sqlStore) CreateCart(ctx context.Context, data *cartmodel.CartCreate) error {
	db := s.db.Table(cartmodel.CartCreate{}.TableName())

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

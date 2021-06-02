package cartstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/cart/cartmodel"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) GetCart(ctx context.Context, userId int, foodId int) (*cartmodel.Cart, error) {
	db := s.db.Table(cartmodel.Cart{}.TableName()).Where("status in (?)", 1)

	var cart cartmodel.Cart
	if err := db.Where(map[string]interface{}{
		"user_id": userId,
		"food_id": foodId,
	}).First(&cart).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &cart, nil
}

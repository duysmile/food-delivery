package orderstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/order/ordermodel"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) GetOrder(ctx context.Context, id int, condition map[string]interface{}) (*ordermodel.Order, error) {
	db := s.db.Table(ordermodel.Order{}.TableName()).Where("status in (?)", 1)

	var data ordermodel.Order
	if condition != nil {
		db = db.Where(condition)
	}
	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}

package orderstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/order/ordermodel"
	"context"
)

func (s *sqlStore) UpdateOrder(ctx context.Context, id int, data *ordermodel.OrderUpdate) error {
	db := s.db.Table(ordermodel.OrderUpdate{}.TableName()).Where("status in (?)", 1)

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

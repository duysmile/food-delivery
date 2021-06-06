package orderstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/order/ordermodel"
	"context"
)

func (s *sqlStore) CreateOrder(ctx context.Context, order *ordermodel.Order, orderDetails []ordermodel.OrderDetail) (int, error) {
	db := s.db.Begin()

	if err := db.Create(order).Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}

	for i := range orderDetails {
		orderDetails[i].OrderId = order.Id
	}

	if err := db.Create(orderDetails).Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}

	return order.Id, nil
}

package ordertrackingstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/ordertracking/ordertrackingmodel"
	"context"
)

func (s *sqlStore) CreateOrderTracking(ctx context.Context, data *ordertrackingmodel.OrderTrackingCreate) error {
	db := s.db.Table(ordertrackingmodel.OrderTrackingCreate{}.TableName())

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

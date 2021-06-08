package ordertrackingstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/ordertracking/ordertrackingmodel"
	"context"
)

func (s *sqlStore) UpdateOrderTracking(ctx context.Context, id int, data *ordertrackingmodel.OrderTrackingUpdate) error {
	db := s.db.Table(ordertrackingmodel.OrderTrackingUpdate{}.TableName()).Where("status in (?)", 1)

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

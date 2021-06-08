package ordertrackingstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/ordertracking/ordertrackingmodel"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) GetOrderTrackingByCondition(ctx context.Context, id int, userId int, condition map[string]interface{}) (*ordertrackingmodel.OrderTracking, error) {
	db := s.db.Table(ordertrackingmodel.OrderTracking{}.TableName()).Where("order_trackings.status in (?)", 1)

	var orderTracking ordertrackingmodel.OrderTracking

	if err := db.Where("order_id = ?", id).
		Joins("JOIN orders ON orders.user_id = ?", userId).
		Where(condition).First(&orderTracking).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &orderTracking, nil
}

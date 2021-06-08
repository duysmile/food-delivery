package ordertrackingstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/ordertracking/ordertrackingmodel"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) GetOrderTrackingByCondition(ctx context.Context, id int, notCondition map[string]interface{}) (*ordertrackingmodel.OrderTracking, error) {
	db := s.db.Table(ordertrackingmodel.OrderTracking{}.TableName()).Where("status in (?)", 1)

	var orderTracking ordertrackingmodel.OrderTracking

	if err := db.Where("id = ?", id).Not(notCondition).First(&orderTracking).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &orderTracking, nil
}

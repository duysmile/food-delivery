package ordertrackingstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/ordertracking/ordertrackingmodel"
	"context"
)

func (s *sqlStore) ListOrderTrackingByCondition(
	ctx context.Context,
	condition map[string]interface{},
	paging *common.Paging,
	moreKeys ...string,
) ([]ordertrackingmodel.OrderTracking, error) {
	db := s.db.Table(ordertrackingmodel.OrderTracking{}.TableName()).Where("status in (?)", 1)

	db = db.Where(condition)
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase58(v); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	db = db.Limit(paging.Limit)

	var orderTrackings []ordertrackingmodel.OrderTracking
	if err := db.Order("id desc").Find(&orderTrackings).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return orderTrackings, nil
}

package orderstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/order/ordermodel"
	"context"
)

func (s *sqlStore) ListOrdersByCondition(
	ctx context.Context,
	condition map[string]interface{},
	paging *common.Paging,
	moreKeys ...string,
) ([]ordermodel.Order, error) {
	db := s.db.Table(ordermodel.Order{}.TableName()).Where("status in (?)", 1)

	db = db.Where(condition)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)
		if err != nil {
			return nil, common.ErrInvalidRequest(err)
		}
		db.Where("id < ?", uid.GetLocalID())
	} else {
		db.Offset((paging.Page - 1) * paging.Limit)
	}

	var orders []ordermodel.Order
	if err := db.Limit(paging.Limit).
		Order("id desc").
		Find(&orders).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return orders, nil
}

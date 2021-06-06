package foodstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/food/foodmodel"
	"context"
)

func (s *sqlStore) ListFoodByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *foodmodel.Filter,
	paging *common.Paging,
	moreInfo ...string,
) ([]foodmodel.Food, error) {
	db := s.db.Table(foodmodel.Food{}.TableName())

	db = db.Where(conditions).Where("status in (?)", 1)

	if f := filter; f != nil {

	}

	if paging != nil {
		if err := db.Where(conditions).Count(&paging.Total).Error; err != nil {
			return nil, common.ErrDB(err)
		}
	}

	for index := range moreInfo {
		db.Preload(moreInfo[index])
	}

	if paging != nil {
		if c := paging.FakeCursor; c != "" {
			if uid, err := common.FromBase58(c); err == nil {
				db = db.Where("id < ?", uid.GetLocalID())
			}
		} else {
			db = db.Offset((paging.Page - 1) * paging.Limit)
		}
		db = db.Limit(paging.Limit)
	}

	var data []foodmodel.Food
	if err := db.
		Order("id desc").
		Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}

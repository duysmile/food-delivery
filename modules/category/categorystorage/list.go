package categorystorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/category/categorymodel"
	"context"
)

func (s *sqlStore) ListCategoryByCondition(
	ctx context.Context,
	condition map[string]interface{},
	paging *common.Paging,
	moreKeys ...string,
) ([]categorymodel.Category, error) {
	db := s.db.Table(categorymodel.Category{}.TableName()).Where("status in (?)", 1)

	db = db.Where(condition)
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase58(v); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	var categories []categorymodel.Category
	if err := db.Order("id desc").
		Limit(paging.Limit).
		Find(&categories).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return categories, nil
}

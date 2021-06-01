package useraddressstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/useraddress/useraddressmodel"
	"context"
)

func (s *sqlStore) ListUserAddresses(
	ctx context.Context,
	userId int,
	paging *common.Paging,
	moreKeys ...string,
) ([]useraddressmodel.UserAddress, error) {
	db := s.db.Table(useraddressmodel.UserAddress{}.TableName()).
		Where("status in (?)", 1)

	db = db.Where("user_id = ?", userId)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase58(paging.FakeCursor); err != nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db = db.Limit(paging.Limit).
			Offset((paging.Page - 1) * paging.Limit)
	}

	var data []useraddressmodel.UserAddress
	if err := db.Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}

package useraddressstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/useraddress/useraddressmodel"
	"context"
)

func (s *sqlStore) GetUserAddress(ctx context.Context, id int, condition map[string]interface{}, moreKeys ...string) (*useraddressmodel.UserAddress, error) {
	db := s.db.Table(useraddressmodel.UserAddress{}.TableName()).
		Where("status in (?)", 1)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var data useraddressmodel.UserAddress
	if err := db.Where("id = ?", id).Where(condition).First(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &data, nil
}

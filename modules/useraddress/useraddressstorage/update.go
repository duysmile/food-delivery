package useraddressstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/useraddress/useraddressmodel"
	"context"
)

func (s *sqlStore) UpdateUserAddress(ctx context.Context, id int, data *useraddressmodel.UserAddressUpdate) error {
	db := s.db.Table(useraddressmodel.UserAddressUpdate{}.TableName()).
		Where("status in (?)", 1)

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

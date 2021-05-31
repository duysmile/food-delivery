package useraddressstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/useraddress/useraddressmodel"
	"context"
)

func (s *sqlStore) CreateUserAddress(ctx context.Context, data *useraddressmodel.UserAddressCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

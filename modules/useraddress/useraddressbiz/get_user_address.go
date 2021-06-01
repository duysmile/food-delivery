package useraddressbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/useraddress/useraddressmodel"
	"context"
)

type GetUserAddressStore interface {
	GetUserAddress(ctx context.Context, id int, condition map[string]interface{}, moreKeys ...string) (*useraddressmodel.UserAddress, error)
}

type getUserAddressBiz struct {
	store GetUserAddressStore
}

func NewGetUserAddressBiz(store GetUserAddressStore) *getUserAddressBiz {
	return &getUserAddressBiz{store: store}
}

func (biz *getUserAddressBiz) GetUserAddress(ctx context.Context, userId int, id int) (*useraddressmodel.UserAddress, error) {
	userAddress, err := biz.store.GetUserAddress(ctx, id, map[string]interface{}{"user_id": userId}, "City")
	if err != nil {
		return nil, common.ErrCannotGetEntity(useraddressmodel.EntityName, err)
	}

	return userAddress, nil
}

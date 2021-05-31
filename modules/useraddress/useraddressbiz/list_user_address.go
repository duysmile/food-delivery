package useraddressbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/useraddress/useraddressmodel"
	"context"
)

type ListUserAddressStore interface {
	ListUserAddresses(
		ctx context.Context,
		userId int,
		paging *common.Paging,
		moreKeys ...string,
	) ([]useraddressmodel.UserAddress, error)
}

type listUserAddressBiz struct {
	store ListUserAddressStore
}

func NewListUserAddressBiz(store ListUserAddressStore) *listUserAddressBiz {
	return &listUserAddressBiz{store: store}
}

func (biz *listUserAddressBiz) ListUserAddresses(
	ctx context.Context,
	userId int,
	paging *common.Paging,
) ([]useraddressmodel.UserAddress, error) {
	store := biz.store

	userAddresses, err := store.ListUserAddresses(ctx, userId, paging, "City")
	if err != nil {
		return nil, common.ErrCannotGetEntity(useraddressmodel.EntityName, err)
	}

	return userAddresses, nil
}

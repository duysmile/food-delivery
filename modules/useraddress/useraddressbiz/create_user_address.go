package useraddressbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/useraddress/useraddressmodel"
	"context"
)

type CreateUserAddressStore interface {
	CreateUserAddress(ctx context.Context, data *useraddressmodel.UserAddressCreate) error
}

type createUserAddressBiz struct {
	store CreateUserAddressStore
}

func NewCreateUserAddressBiz(store CreateUserAddressStore) *createUserAddressBiz {
	return &createUserAddressBiz{store: store}
}

func (biz *createUserAddressBiz) CreateUserAddress(
	ctx context.Context,
	data *useraddressmodel.UserAddressCreate,
) error {
	store := biz.store

	if err := store.CreateUserAddress(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(useraddressmodel.EntityName, err)
	}

	return nil

}

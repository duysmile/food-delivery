package useraddressbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/useraddress/useraddressmodel"
	"context"
)

type UpdateUserAddressStore interface {
	UpdateUserAddress(ctx context.Context, id int, data *useraddressmodel.UserAddressUpdate) error
	GetUserAddress(ctx context.Context, id int, condition map[string]interface{}, moreKeys ...string) (*useraddressmodel.UserAddress, error)
}

type updateUserAddressBiz struct {
	store UpdateUserAddressStore
}

func NewUpdateUserAddressBiz(store UpdateUserAddressStore) *updateUserAddressBiz {
	return &updateUserAddressBiz{store: store}
}

func (biz *updateUserAddressBiz) UpdateUserAddress(
	ctx context.Context,
	userId int,
	id int,
	data *useraddressmodel.UserAddressUpdate,
) error {
	_, err := biz.store.GetUserAddress(ctx, id, map[string]interface{}{"user_id": userId})

	if err != nil {
		return common.ErrCannotUpdateEntity(useraddressmodel.EntityName, err)
	}

	if err = biz.store.UpdateUserAddress(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(useraddressmodel.EntityName, err)
	}

	return nil
}

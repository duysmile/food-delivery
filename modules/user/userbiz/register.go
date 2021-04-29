package userbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/user/usermodel"
	"context"
)

type RegisterStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBiz struct {
	store  RegisterStore
	hasher Hasher
}

func NewRegisterBiz(store RegisterStore, hasher Hasher) *registerBiz {
	return &registerBiz{
		store:  store,
		hasher: hasher,
	}
}

func (biz *registerBiz) Register(ctx context.Context, data *usermodel.UserCreate) error {
	if user, _ := biz.store.FindUser(
		ctx,
		map[string]interface{}{
			"email": data.Email,
		},
	); user != nil {
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}

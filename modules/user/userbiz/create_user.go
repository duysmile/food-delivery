package userbiz

import (
	"200lab/food-delivery/modules/user/usermodel"
	"context"
)

type CreateUserStore interface {
	Create(ctx context.Context, data *usermodel.UserCreate) error
}

type createUserBiz struct {
	store CreateUserStore
}

func NewCreateUserBiz(store CreateUserStore) *createUserBiz {
	return &createUserBiz{store: store}
}

func (c *createUserBiz) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	err := c.store.Create(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

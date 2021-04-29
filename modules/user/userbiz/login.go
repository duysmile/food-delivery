package userbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component/tokenprovider"
	"200lab/food-delivery/modules/user/usermodel"
	"context"
)

type LoginStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBiz struct {
	store         LoginStore
	hasher        Hasher
	tokenProvider tokenprovider.Provider
	expiry        int
}

func NewLoginBiz(
	store LoginStore,
	hasher Hasher,
	tokenProvider tokenprovider.Provider,
	expiry int,
) *loginBiz {
	return &loginBiz{
		store:         store,
		hasher:        hasher,
		tokenProvider: tokenProvider,
		expiry:        expiry,
	}
}

// 1. Find user, email
// 2. Hash pass from input and compare with pass in db
// 3. Provider: issue JWT token for client
// 3.1. Access token and refresh token
// 4. Return token(s)

func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	passHashed := biz.hasher.Hash(data.Password + user.Salt)
	if user.Password != passHashed {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserID: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}

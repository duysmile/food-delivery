package jwt

import (
	"200lab/food-delivery/component/tokenprovider"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtProvider struct {
	secret string
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func NewTokenProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	return &tokenprovider.Token{
		Token:   myToken,
		Expiry:  expiry,
		Created: time.Now(),
	}, nil
}

func (j *jwtProvider) Validate(token string) (*tokenprovider.TokenPayload, error) {
	return nil, nil
}

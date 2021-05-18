package component

import (
	"200lab/food-delivery/component/uploadprovider"
	"200lab/food-delivery/pubsub"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetUploadProvider() uploadprovider.UploadProvider
	GetSecretKey() string
	GetPubSub() pubsub.Pubsub
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
	pubsub         pubsub.Pubsub
}

func NewAppContext(
	db *gorm.DB,
	uploadProvider uploadprovider.UploadProvider,
	secretKey string,
	pb pubsub.Pubsub,
) *appCtx {
	return &appCtx{db: db, uploadProvider: uploadProvider, secretKey: secretKey, pubsub: pb}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetUploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}

func (ctx *appCtx) GetSecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetPubSub() pubsub.Pubsub {
	return ctx.pubsub
}

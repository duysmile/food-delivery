package ginuser

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/component/hasher"
	"200lab/food-delivery/component/tokenprovider/jwt"
	"200lab/food-delivery/modules/user/userbiz"
	"200lab/food-delivery/modules/user/usermodel"
	"200lab/food-delivery/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())
		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewLoginBiz(store, md5, tokenProvider, 60*60*24*30)
		account, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}

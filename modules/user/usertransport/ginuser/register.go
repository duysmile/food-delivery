package ginuser

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/component/hasher"
	"200lab/food-delivery/modules/user/userbiz"
	"200lab/food-delivery/modules/user/usermodel"
	"200lab/food-delivery/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data usermodel.UserCreate

		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBiz(store, md5)

		if err := biz.Register(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}

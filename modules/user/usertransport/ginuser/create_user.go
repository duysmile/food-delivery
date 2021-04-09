package ginuser

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/user/userbiz"
	"200lab/food-delivery/modules/user/usermodel"
	"200lab/food-delivery/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data usermodel.UserCreate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userbiz.NewCreateUserBiz(store)

		if err := biz.CreateUser(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
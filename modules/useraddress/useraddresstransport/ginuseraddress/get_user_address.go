package ginuseraddress

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/useraddress/useraddressbiz"
	"200lab/food-delivery/modules/useraddress/useraddressstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserAddress(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}

		store := useraddressstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := useraddressbiz.NewGetUserAddressBiz(store)

		userAddress, err := biz.GetUserAddress(c.Request.Context(), user.GetUserId(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		userAddress.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(userAddress))
	}
}

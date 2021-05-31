package ginuseraddress

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/useraddress/useraddressbiz"
	"200lab/food-delivery/modules/useraddress/useraddressmodel"
	"200lab/food-delivery/modules/useraddress/useraddressstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserAddress(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		var userAddress useraddressmodel.UserAddressCreate
		if err := c.ShouldBind(&userAddress); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		userAddress.UserId = user.GetUserId()
		store := useraddressstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := useraddressbiz.CreateUserAddressStore(store)

		if err := biz.CreateUserAddress(c.Request.Context(), &userAddress); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

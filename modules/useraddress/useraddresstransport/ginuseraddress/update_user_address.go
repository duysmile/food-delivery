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

func UpdateUserAddress(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data useraddressmodel.UserAddressUpdate
		if err = c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := useraddressstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := useraddressbiz.NewUpdateUserAddressBiz(store)

		if err = biz.UpdateUserAddress(
			c.Request.Context(),
			user.GetUserId(),
			int(uid.GetLocalID()),
			&data,
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

package ginuseraddress

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/useraddress/useraddressbiz"
	"200lab/food-delivery/modules/useraddress/useraddressstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUserAddresses(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := useraddressstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := useraddressbiz.NewListUserAddressBiz(store)

		userAddresses, err := biz.ListUserAddresses(c.Request.Context(), user.GetUserId(), &paging)
		if err != nil {
			panic(err)
		}

		for i := range userAddresses {
			userAddresses[i].Mask(false)
			if i == len(userAddresses)-1 {
				paging.NextCursor = userAddresses[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(userAddresses, paging, nil))
	}
}

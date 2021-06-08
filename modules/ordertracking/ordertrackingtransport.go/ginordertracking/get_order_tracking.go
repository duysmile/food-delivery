package ginordertracking

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/ordertracking/ordertrackingbiz"
	"200lab/food-delivery/modules/ordertracking/ordertrackingstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrderTracking(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := ordertrackingstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := ordertrackingbiz.NewGetOrderTrackingBiz(store)

		orderTracking, err := biz.GetOrderTracking(c.Request.Context(), user.GetUserId(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		orderTracking.GenUID(common.DbTypeOrderTracking)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(orderTracking))
	}
}

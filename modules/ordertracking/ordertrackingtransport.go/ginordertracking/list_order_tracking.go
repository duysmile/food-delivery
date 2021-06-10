package ginordertracking

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/order/orderstorage"
	"200lab/food-delivery/modules/ordertracking/ordertrackingbiz"
	"200lab/food-delivery/modules/ordertracking/ordertrackingstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOrderTracking(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := ordertrackingstorage.NewSQLStore(appCtx.GetMainDBConnection())
		orderStore := orderstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := ordertrackingbiz.NewListOrderTrackingBiz(store, orderStore)

		orderTrackings, err := biz.ListOrderTracking(c.Request.Context(), user.GetUserId(), int(uid.GetLocalID()), &paging)
		if err != nil {
			panic(err)
		}

		for i := range orderTrackings {
			orderTrackings[i].GenUID(common.DbTypeOrderTracking)
			if i == len(orderTrackings)-1 {
				paging.NextCursor = orderTrackings[i].FakeId.String()
			}
		}
		// orderTracking.GenUID(common.DbTypeOrderTracking)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(orderTrackings))
	}
}

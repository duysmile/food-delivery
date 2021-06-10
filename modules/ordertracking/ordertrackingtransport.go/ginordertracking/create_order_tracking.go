package ginordertracking

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/order/orderstorage"
	"200lab/food-delivery/modules/ordertracking/ordertrackingbiz"
	"200lab/food-delivery/modules/ordertracking/ordertrackingmodel"
	"200lab/food-delivery/modules/ordertracking/ordertrackingstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrderTracking(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		var data ordertrackingmodel.OrderTrackingCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if uid, err := common.FromBase58(data.FakeOrderId.String()); err != nil {
			panic(common.ErrInvalidRequest(err))
		} else {
			data.OrderId = int(uid.GetLocalID())
		}

		store := ordertrackingstorage.NewSQLStore(appCtx.GetMainDBConnection())
		orderStore := orderstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := ordertrackingbiz.NewCreateOrderBiz(store, orderStore)

		if err := biz.CreateOrderTracking(c.Request.Context(), user.GetUserId(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

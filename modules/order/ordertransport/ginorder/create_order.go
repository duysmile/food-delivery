package ginorder

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/food/foodstorage"
	"200lab/food-delivery/modules/order/orderbiz"
	"200lab/food-delivery/modules/order/ordermodel"
	"200lab/food-delivery/modules/order/orderstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		var data ordermodel.OrderCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := data.Validate(); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data.Unmask()

		store := orderstorage.NewSQLStore(appCtx.GetMainDBConnection())
		foodStore := foodstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := orderbiz.NewCreateOrderBiz(store, foodStore, appCtx.GetPubSub())

		if err := biz.CreateOrder(c.Request.Context(), user.GetUserId(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

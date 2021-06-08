package ginorder

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/order/orderbiz"
	"200lab/food-delivery/modules/order/orderstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOrder(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := orderstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := orderbiz.NewListOrderBiz(store)

		orders, err := biz.ListOrder(c.Request.Context(), user.GetUserId(), &paging)
		if err != nil {
			panic(err)
		}

		for i := range orders {
			orders[i].GenUID(common.DbTypeOrder)
			if i == len(orders)-1 {
				paging.NextCursor = orders[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(orders, paging, nil))
	}
}

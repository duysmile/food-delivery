package gincart

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/cart/cartbiz"
	"200lab/food-delivery/modules/cart/cartmodel"
	"200lab/food-delivery/modules/cart/cartstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCart(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		var cart cartmodel.CartCreate
		if err := c.ShouldBind(&cart); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		cart.UserId = user.GetUserId()
		cart.FoodId = int(cart.FakeFoodId.GetLocalID())

		store := cartstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := cartbiz.NewCreateCartBiz(store)

		if err := biz.CreateCart(c.Request.Context(), &cart); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

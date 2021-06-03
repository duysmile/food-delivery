package gincart

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/cart/cartbiz"
	"200lab/food-delivery/modules/cart/cartstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCart(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := cartstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := cartbiz.NewGetCartBiz(store)

		cart, err := biz.ListItemInCartByUserId(c.Request.Context(), user.GetUserId(), &paging)
		if err != nil {
			panic(err)
		}

		for i := range cart {
			cart[i].Food.Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(cart, paging, nil))
	}
}

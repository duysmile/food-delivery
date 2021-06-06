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

func UpdateCart(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data cartmodel.CartUpdate
		if err = c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := cartstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := cartbiz.NewUpdateCartBiz(store)

		if err = biz.UpdateCart(c.Request.Context(), user.GetUserId(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

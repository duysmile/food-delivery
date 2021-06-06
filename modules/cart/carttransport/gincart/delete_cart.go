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

func DeleteCart(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data := cartmodel.CartDelete{
			UserId: user.GetUserId(),
			FoodId: int(uid.GetLocalID()),
		}

		store := cartstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := cartbiz.NewDeleteCartBiz(store)

		if err = biz.DeleteCart(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

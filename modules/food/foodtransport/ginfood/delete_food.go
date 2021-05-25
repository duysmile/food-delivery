package ginfood

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/food/foodbiz"
	"200lab/food-delivery/modules/food/foodstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteFood(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)
		foodId, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := foodstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := foodbiz.NewDeleteFoodBiz(store)
		if err := biz.DeleteFood(c.Request.Context(), user.GetUserId(), int(foodId.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

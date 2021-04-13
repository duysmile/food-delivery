package ginrestaurant

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/restaurant/restaurantbiz"
	"200lab/food-delivery/modules/restaurant/restaurantmodel"
	"200lab/food-delivery/modules/restaurant/restaurantstorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data restaurantmodel.RestaurantUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)
		if err := biz.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

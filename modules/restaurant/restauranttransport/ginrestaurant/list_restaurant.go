package ginrestaurant

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/restaurant/restaurantbiz"
	restarantmodel "200lab/food-delivery/modules/restaurant/restaurantmodel"
	"200lab/food-delivery/modules/restaurant/restaurantstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRestaurantByCondition(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restarantmodel.Filter
		var paging common.Paging
		var result []restarantmodel.Restaurant

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		paging.Fulfill()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListUserBiz(store)

		result, err := biz.ListUserBiz(c.Request.Context(), nil, &filter, &paging)

		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}

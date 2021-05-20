package ginrestaurant

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/restaurant/restaurantbiz"
	restaurantmodel "200lab/food-delivery/modules/restaurant/restaurantmodel"
	"200lab/food-delivery/modules/restaurant/restaurantstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRestaurantByCondition(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter
		var paging common.Paging
		var result []restaurantmodel.Restaurant

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		// likeStore := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		// biz := restaurantbiz.NewListUserBiz(store, likeStore)
		biz := restaurantbiz.NewListUserBiz(store)

		result, err := biz.ListUserBiz(c.Request.Context(), nil, &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

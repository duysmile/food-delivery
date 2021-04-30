package ginrestaurant

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/restaurant/restaurantbiz"
	restaurantmodel "200lab/food-delivery/modules/restaurant/restaurantmodel"
	"200lab/food-delivery/modules/restaurant/restaurantstorage"
	"200lab/food-delivery/modules/upload/uploadstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		user := c.MustGet(common.CurrentUser).(common.Requester)

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data.OwnerId = user.GetUserId()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		imageStore := uploadstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store, imageStore)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.GenUID(common.DbTypeRestaurant)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}

}

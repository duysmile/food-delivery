package ginfood

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/food/foodbiz"
	"200lab/food-delivery/modules/food/foodmodel"
	"200lab/food-delivery/modules/food/foodstorage"
	"200lab/food-delivery/modules/restaurant/restaurantstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFood(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		var food *foodmodel.FoodCreate

		if err := c.ShouldBind(&food); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := food.Unmask(); err != nil {
			panic(err)
		}

		store := foodstorage.NewSQLStore(appCtx.GetMainDBConnection())
		restaurantStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := foodbiz.NewCreateFoodBiz(store, restaurantStore)

		if err := biz.CreateFood(c.Request.Context(), user.GetUserId(), food); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

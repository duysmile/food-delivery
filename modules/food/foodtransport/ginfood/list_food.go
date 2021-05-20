package ginfood

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/food/foodbiz"
	"200lab/food-delivery/modules/food/foodstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListFoodByRestaurantId(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		restaurantId := int(uid.GetLocalID())
		store := foodstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := foodbiz.NewListFoodBiz(store)

		foods, err := biz.ListFoodByRestaurantId(c.Request.Context(), restaurantId, &paging)
		if err != nil {
			panic(err)
		}

		for i := range foods {
			foods[i].Mask(false)
			if i == len(foods)-1 {
				paging.NextCursor = foods[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(foods, paging, nil))
	}
}

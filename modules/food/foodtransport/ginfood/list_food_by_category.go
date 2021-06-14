package ginfood

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/category/categorystorage"
	"200lab/food-delivery/modules/food/foodbiz"
	"200lab/food-delivery/modules/food/foodstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListFoodByCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err = c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := foodstorage.NewSQLStore(appCtx.GetMainDBConnection())
		categoryStore := categorystorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := foodbiz.NewListFoodByCategoryBiz(store, categoryStore)

		foods, err := biz.ListFoodByCategory(c.Request.Context(), int(uid.GetLocalID()), &paging)
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

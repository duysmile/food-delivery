package gincategory

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/category/categorybiz"
	"200lab/food-delivery/modules/category/categorystorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := categorystorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewListCategoryBiz(store)

		categories, err := biz.ListCategory(c.Request.Context(), &paging)
		if err != nil {
			panic(err)
		}

		for i := range categories {
			categories[i].GenUID(common.DbTypeCategory)

			if i == len(categories)-1 {
				paging.NextCursor = categories[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(categories, paging, nil))
	}
}

package gincategory

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/category/categorybiz"
	"200lab/food-delivery/modules/category/categorymodel"
	"200lab/food-delivery/modules/category/categorystorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var category categorymodel.CategoryCreate
		if err := c.ShouldBind(&category); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewCreateCategoryBiz(store)

		if err := biz.CreateCategoryBiz(c.Request.Context(), &category); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

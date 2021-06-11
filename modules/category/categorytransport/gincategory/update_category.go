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

func UpdateCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var dataUpdate categorymodel.CategoryUpdate
		if err = c.ShouldBind(&dataUpdate); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewUpdateCategoryBiz(store)

		if err = biz.UpdateCategory(c.Request.Context(), int(uid.GetLocalID()), &dataUpdate); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

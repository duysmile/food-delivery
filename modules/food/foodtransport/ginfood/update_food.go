package ginfood

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/food/foodbiz"
	"200lab/food-delivery/modules/food/foodmodel"
	"200lab/food-delivery/modules/food/foodstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateFood(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)
		id, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var foodData foodmodel.FoodUpdate
		if err = c.ShouldBind(&foodData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if foodData.FakeCategoryId != "" {
			if uid, err := common.FromBase58(foodData.FakeCategoryId); err != nil {
				panic(common.ErrInvalidRequest(err))
			} else {
				categoryId := int(uid.GetLocalID())
				foodData.CategoryId = &categoryId
			}
		}

		store := foodstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := foodbiz.NewUpdateFoodBiz(store)

		if err := biz.UpdateFood(c.Request.Context(), user.GetUserId(), int(id.GetLocalID()), &foodData); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

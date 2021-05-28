package ginfoodlike

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/foodlike/foodlikebiz"
	"200lab/food-delivery/modules/foodlike/foodlikemodel"
	"200lab/food-delivery/modules/foodlike/foodlikestorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UnlikeFood(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)
		foodId, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := foodlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := foodlikebiz.NewUnlikeFoodBiz(store, appCtx.GetPubSub())
		if err := biz.UnlikeFood(c.Request.Context(), &foodlikemodel.FoodLikeCreate{
			UserId: user.GetUserId(),
			FoodId: int(foodId.GetLocalID()),
		}); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

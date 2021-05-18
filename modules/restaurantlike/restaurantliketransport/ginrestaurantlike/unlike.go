package ginrestaurantlike

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/restaurant/restaurantstorage"
	"200lab/food-delivery/modules/restaurantlike/restaurantlikebiz"
	"200lab/food-delivery/modules/restaurantlike/restaurantlikemodel"
	"200lab/food-delivery/modules/restaurantlike/restaurantlikestorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UnlikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)
		restaurantId, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		likeDelete := restaurantlikemodel.LikeDelete{
			RestaurantId: int(restaurantId.GetLocalID()),
			UserId:       user.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		decreaseLikedCountStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantlikebiz.NewDeleteRestaurantLikeBiz(store, decreaseLikedCountStore)

		if err := biz.DeleteLike(c.Request.Context(), &likeDelete); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

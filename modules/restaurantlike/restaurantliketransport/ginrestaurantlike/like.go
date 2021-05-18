package ginrestaurantlike

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/restaurantlike/restaurantlikebiz"
	"200lab/food-delivery/modules/restaurantlike/restaurantlikemodel"
	"200lab/food-delivery/modules/restaurantlike/restaurantlikestorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)
		restaurantId, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		restaurantLike := restaurantlikemodel.LikeCreate{
			UserId:       user.GetUserId(),
			RestaurantId: int(restaurantId.GetLocalID()),
		}

		createStore := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		// increaseLikedCountStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantlikebiz.NewCreateRestaurantLikeBiz(createStore, appCtx.GetPubSub())

		if err := biz.CreateLike(c.Request.Context(), &restaurantLike); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

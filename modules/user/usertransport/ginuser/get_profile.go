package ginuser

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfileUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

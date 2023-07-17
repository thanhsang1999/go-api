package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"go-api/common"
	"go-api/component/appctx"
	restaurantbusiness "go-api/module/restaurant/business"
	restaurantstorage "go-api/module/restaurant/storage"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

		business := restaurantbusiness.NewDeleteRestaurantBusiness(store)
		if err := business.DeleteRestaurant(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

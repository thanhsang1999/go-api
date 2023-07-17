package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"go-api/common"
	"go-api/component/appctx"
	restaurantbusiness "go-api/module/restaurant/business"
	restaurantmodel "go-api/module/restaurant/model"
	restaurantstorage "go-api/module/restaurant/storage"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		business := restaurantbusiness.NewCreateRestaurantBusiness(store)
		err := business.CreateRestaurant(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}

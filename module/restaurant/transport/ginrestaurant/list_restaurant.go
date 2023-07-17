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

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		pagingData.Fulfill()
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		business := restaurantbusiness.NewListRestaurantBusiness(store)

		result, err := business.ListRestaurant(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}

package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"go-api/common"
	"go-api/component/appctx"
	restaurantbusiness "go-api/module/restaurant/business"
	restaurantstorage "go-api/module/restaurant/storage"
	"net/http"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("id"))

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

		business := restaurantbusiness.NewDeleteRestaurantBusiness(store)
		if err := business.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

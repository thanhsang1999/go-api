package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"go-api/common"
	"go-api/component/appctx"
	restaurantbusiness "go-api/module/restaurant/business"
	restaurantmodel "go-api/module/restaurant/model"
	restaurantstorage "go-api/module/restaurant/storage"
	"go-api/module/user/usermodel"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		//go func() {
		//	defer common.AppRecover()
		//	arr := []int{}
		//	log.Println(arr[0])
		//}()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		business := restaurantbusiness.NewCreateRestaurantBusiness(store)
		user := c.MustGet(common.CurrentUser).(usermodel.User)
		data.OwnerId = user.Id
		err := business.CreateRestaurant(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}

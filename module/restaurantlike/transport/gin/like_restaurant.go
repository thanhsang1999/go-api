package ginrestaurantlike

import (
	"github.com/gin-gonic/gin"
	"go-api/common"
	"go-api/component/appctx"
	restaurantlikebiz "go-api/module/restaurantlike/biz"
	restaurantlikemodel "go-api/module/restaurantlike/model"
	restaurantlikestorage "go-api/module/restaurantlike/storage"
	"go-api/module/user/usermodel"
	"net/http"
)

func LikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		business := restaurantlikebiz.NewCreateRestaurantLikeBiz(store, appCtx.GetPubSub())
		user := c.MustGet(common.CurrentUser).(usermodel.User)

		var data = &restaurantlikemodel.RestaurantLikes{
			RestaurantId: int(id.GetLocalID()),
			UserId:       user.Id,
		}
		if err := business.CreateRestaurantLike(c.Request.Context(), data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

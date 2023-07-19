package middleware

import (
	"github.com/gin-gonic/gin"
	"go-api/common"
	"go-api/component/appctx"
	"go-api/module/user/usermodel"
)

func Authenticate(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := usermodel.User{
			Id:    99,
			Email: "abc@gmail.com",
		}
		c.Set(common.CurrentUser, user)
		c.Next()
	}
}

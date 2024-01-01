package middleware

import (
	auth_response "go-project/data/responses/auth"
	"go-project/helper"
	"go-project/initializers"
	"go-project/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "UnAuthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func RoleMiddleware(roles string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sep string = ","
		rolesList := strings.Split(roles, sep)
		currentUser := &auth_response.AuthResponse{}
		err := utils.TokenValid(ctx)
		if err != nil {
			helper.Error(ctx, err)
			ctx.Abort()
			return
		}

		user_id, err := utils.ExtractTokenID(ctx)
		initializers.DB.Table("users").Where("id", user_id).First(currentUser)
		if err != nil {
			helper.Error(ctx, err)
			ctx.Abort()
			return
		}
		exist, _ := helper.InArray(currentUser.Role, rolesList)

		if !exist {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "You have not a right permission"})
			ctx.Abort()
			return
		}

		ctx.Next()

	}
}

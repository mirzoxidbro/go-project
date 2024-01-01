package controller

import (
	"go-project/data/requests/auth"
	auth_response "go-project/data/responses/auth"
	"go-project/helper"
	"go-project/initializers"
	"go-project/model"
	"go-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var input auth.LoginRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := model.User{}

	u.Email = input.Email
	u.Password = input.Password

	token, err := model.LoginCheck(u.Email, u.Password)

	if err != nil {
		helper.Error(ctx, err)
		return
	}

	response := helper.RespondWithJSON(token)
	ctx.JSON(http.StatusBadRequest, response)
}

func CurrentUser(ctx *gin.Context) {
	currentUser := &auth_response.AuthResponse{}
	err := utils.TokenValid(ctx)
	if err != nil {
		response := helper.RespondWithJSON(err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user_id, err := utils.ExtractTokenID(ctx)
	initializers.DB.Table("users").Where("id", user_id).First(currentUser)

	if err != nil {
		helper.Error(ctx, err)
		return
	}
	response := helper.RespondWithJSON(currentUser)
	ctx.JSON(http.StatusOK, response)
}

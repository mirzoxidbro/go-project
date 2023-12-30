package helper

import (
	response "go-project/data/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error(ctx *gin.Context, err error) {
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusBadRequest,
			Status: "Error",
			Data:   err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
	}
}

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func RespondWithJSON(payload interface{}) response.Response {
	Response := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   payload,
	}

	return Response
}

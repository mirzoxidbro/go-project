package controllers

import (
	"net/http"

	"github.com/mirzoxidbro/go-project/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
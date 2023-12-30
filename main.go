package main

import (
	"go-project/helper"
	"go-project/migrate"
	"go-project/router"
	"net/http"
)

func main() {

	migrate.RunMigration()

	routes := router.NewRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}

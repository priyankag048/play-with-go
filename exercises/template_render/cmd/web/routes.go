package main

import (
	"net/http"
	"template_render/pkg/config"
	"template_render/pkg/handlers"

	"github.com/gorilla/mux"
)

func getHandler(app *config.AppConfig) http.Handler {
	router := mux.NewRouter()
	router.Use(checkCSRF)
	router.Use(sessionLoad)
	router.HandleFunc("/", handlers.Repo.Home).Methods("GET")
	router.HandleFunc("/about", handlers.Repo.About).Methods("GET")
	return router
}

package router

import (
	"netflix-backend/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.CreateMovie)

	return router
}

package routes

import (
	"github.com/fdiezdev/edcomments/controllers"
	"github.com/gorilla/mux"
)

// SetLoginRouter -> login router
func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login/", controllers.Login).Methods("POST")
}

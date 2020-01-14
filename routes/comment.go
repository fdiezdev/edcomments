package routes

import (
	"github.com/fdiezdev/edcomments/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetCommentRouter -> establish comment creation route
func SetCommentRouter(router *mux.Router) {
	prefix := "/api/comments"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)

	/*
		/api/comments/ => POST
	*/
	subRouter.HandleFunc("/", controllers.CreateComment).Methods("POST")
	/*
		/api/comments/ => GET
	*/
	subRouter.HandleFunc("/", controllers.GetComments).Methods("GET")
	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)
}

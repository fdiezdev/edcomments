package routes

import (
	"github.com/fdiezdev/edcomments/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetVoteRouter -> router to create a vote
func SetVoteRouter(router *mux.Router) {
	prefix := "/api/votes"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)

	subRouter.HandleFunc("/", controllers.VoteRegister).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)

}

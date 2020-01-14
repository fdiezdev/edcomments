package routes

import (
	"github.com/gorilla/mux"
)

// InitRoutes -> starts routes, returns router
func InitRoutes() *mux.Router {

	router := mux.NewRouter().StrictSlash(false)
	SetLoginRouter(router)
	SetUserRouter(router)
	SetCommentRouter(router)
	SetVoteRouter(router)

	return router
}

package server

import (
	getusers "github.com/MSex/lab-go-http/app/server/endpoints/users/get"
	listusers "github.com/MSex/lab-go-http/app/server/endpoints/users/list"
	postusers "github.com/MSex/lab-go-http/app/server/endpoints/users/post"
	"github.com/julienschmidt/httprouter"
)

// ProvideRouter builds and return router
func ProvideRouter(
	listUsersHandler *listusers.Handler,
	getUsersHandler *getusers.Handler,
	postUsersHandler *postusers.Handler,
) *httprouter.Router {
	router := httprouter.New()
	//todo inject handlers
	// router.GET("/", Index)
	// router.GET("/hello/:userId", Hello)
	router.GET("/users", listUsersHandler.Handle)
	router.POST("/users", postUsersHandler.Handle)
	router.GET("/users/:userId", getUsersHandler.Handle)

	return router
}

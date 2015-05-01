package main

import (
	"os"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"github.com/justphil/httprouter-controller"

	"github.com/justphil/denatify-service/controllers"
)

func main() {
	// init the main router
	router := httprouter.New()

	// add routes
	router.GET("/", controller.Action((*controllers.LandingPageController).Index))
	router.GET("/users", controller.Action((*controllers.UsersController).Users))
	router.GET("/user/new", controller.Action((*controllers.UsersController).NewUser))
	router.POST("/user/new", controller.Action((*controllers.UsersController).CreateNewUser))
	router.GET("/users/:userId", controller.Action((*controllers.UsersController).User))
	router.GET("/users/:userId/delete", controller.Action((*controllers.UsersController).DeleteUser))
	router.GET("/users/:userId/perform-deletion", controller.Action((*controllers.UsersController).PerformUserDeletion))
	router.GET("/users/:userId/update", controller.Action((*controllers.UsersController).UpdateUser))
	router.POST("/users/:userId/perform-update", controller.Action((*controllers.UsersController).PerformUserUpdate))

	// initialize middlewares (includes negroni.Recovery, negroni.Logging, negroni.Static)
	n := negroni.Classic()

	// the main router should always be the last handler
	n.UseHandler(router)

	// bind server to interface / port and block
	n.Run(":" + getPort())
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return port
}

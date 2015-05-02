package main

import (
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justphil/httprouter-controller"

	"github.com/justphil/denatify-service/controllers"
	"github.com/justphil/denatify-service/middlewares"
)

func main() {
	// init router
	r := httprouter.New()

	// add routes
	r.GET("/", controller.Action((*controllers.LandingPageController).Index))
	r.GET("/users", controller.Action((*controllers.UsersController).Users))
	r.GET("/user/new", controller.Action((*controllers.UsersController).NewUser))
	r.POST("/user/new", controller.Action((*controllers.UsersController).CreateNewUser))
	r.GET("/users/:userId", controller.Action((*controllers.UsersController).User))
	r.GET("/users/:userId/delete", controller.Action((*controllers.UsersController).DeleteUser))
	// TODO: handle this by using HTTP DELETE
	r.GET("/users/:userId/perform-deletion", controller.Action((*controllers.UsersController).PerformUserDeletion))
	r.GET("/users/:userId/update", controller.Action((*controllers.UsersController).UpdateUser))
	// TODO: handle this by using HTTP PUT
	r.POST("/users/:userId/perform-update", controller.Action((*controllers.UsersController).PerformUserUpdate))

	// init negroni and add default middlewares (negroni.Recovery, negroni.Logging, negroni.Static)
	n := negroni.Classic()

	// add more middlewares
	n.Use(middlewares.FooMiddleware())

	// the router should always be the last handler
	n.UseHandler(context.ClearHandler(r)) // ClearHandler clears context values at the end of a request lifetime

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

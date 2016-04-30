package main


import (
	"net/http"

	"MgoMuxTest/dbRest/controllers"
)

type Route struct {
	Name 		string
	Method		string
	Pattern		string
	HandlerFunc	http.HandlerFunc
}

type Routes []Route

// pass in db connection to controller
var userController = controllers.NewUserController(getConnection())

var routes = Routes {

	Route {
		"Index",
		"GET",
		"/",
		Index,
	},
	Route {
		"AllUser",
		"GET",
		"/users",
		userController.GetAllUser,
	},
	
	Route {
		"FindUser",
		"GET",
		"/users/{userId}",
		userController.FindUser,
	},
	Route {
		"CreateUser",
		"POST",
		"/user",
		userController.CreateUser,
	},
}
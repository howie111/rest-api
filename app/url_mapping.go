package app

import (
	. "github.com/howie111/rest-api/controllers/ping"
	. "github.com/howie111/rest-api/controllers/users"
)

func UrlMapping() {

	router.GET("/ping", PingController)
	router.POST("/users", CreateUser)
	router.GET("/users/:user_id", GetUser)
	//router.GET("/users/search", controllers.SearchUser)
}

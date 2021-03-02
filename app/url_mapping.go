package app

import (
	. "github.com/howie111/rest-api/controllers/ping"
	. "github.com/howie111/rest-api/controllers/users"
)

func UrlMapping() {

	router.GET("/ping", PingController)
	router.POST("/users", Create)
	router.GET("/users/:user_id", Get)
	router.PUT("/users/:user_id", Update)
	router.PATCH("/users/:user_id", Update)
	router.DELETE("/users/:user_id", Delete)
	router.GET("internal/users/search", Search)
}

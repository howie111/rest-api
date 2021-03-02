package app

import (
	"github.com/gin-gonic/gin"
	"github.com/howie111/rest-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {

	logger.Info("about to start the application")

	UrlMapping()
	router.Run(":8080")
}

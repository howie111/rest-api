package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingController(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

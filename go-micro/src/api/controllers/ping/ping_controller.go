package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ping = "ping"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, ping)
}

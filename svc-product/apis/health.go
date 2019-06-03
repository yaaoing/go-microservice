package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ConsulCheckApi(c *gin.Context) {
	c.String(http.StatusOK, "consulCheck")
}

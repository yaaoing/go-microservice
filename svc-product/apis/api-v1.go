package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var log = logrus.New()

func CreateAccount(c *gin.Context) {
	client := registry.GetClient()
	agentService, _, err := client.Agent().Service("account-service", nil)
	if err != nil {
		log.Error("failed to fetch service")
		c.String(http.StatusInternalServerError, "error occur")
	}
	address := agentService.Address

}

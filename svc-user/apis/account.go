package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"leo/go-microservice/svc-user/models"
	"net/http"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddPersionApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	account := models.Account{FirstName: firstName, LastName: lastName}

	ra, err := account.CreateAccount()
	if err != nil {
		log.Info("failed to create a new account")
	}
	msg := fmt.Sprintf("create succeddful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

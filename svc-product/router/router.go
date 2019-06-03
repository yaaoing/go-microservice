package router

import (
	"github.com/gin-gonic/gin"
	"leo/go-microservice/svc-user/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/check", apis.ConsulCheckApi)

	router.GET("/api/v1/account", apis.IndexApi)

	return router
}

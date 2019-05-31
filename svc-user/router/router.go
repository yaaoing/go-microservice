package router

import (
	"github.com/gin-gonic/gin"
	"leo/go-microservice/svc-user/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/check", apis.ConsulCheckApi)

	router.GET("/", apis.IndexApi)
	router.POST("/account", apis.AddPersionApi)

	return router
}

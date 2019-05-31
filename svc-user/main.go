package main

import (
	db "leo/go-microservice/svc-user/databases"
	"leo/go-microservice/svc-user/registry"
	"leo/go-microservice/svc-user/router"
)

func main() {
	defer db.DisConnect()

	registry.RegistryServer()
	router := router.InitRouter()
	router.Run(":8080")
}

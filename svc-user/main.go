package main

import (
	db "leo/go-microservice/svc-user/databases"
	"leo/go-microservice/svc-user/registry"
	"leo/go-microservice/svc-user/router"
)

func main() {
	defer db.DisConnect()

	reg := new(registry.ZookeeperClient)
	reg.RegistryServer()

	router := router.InitRouter()
	router.Run(":8080")
}

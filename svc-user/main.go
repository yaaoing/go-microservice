package main

import (
	db "leo/go-microservice/svc-user/databases"
	"leo/go-microservice/svc-user/router"
)

func main() {
	defer db.SqlDB.Close()
}

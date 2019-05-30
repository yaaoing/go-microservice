package main

import (
	"fmt"
	"log"

	"net/http"

	consulapi "github.com/hashicorp/consul/api"
)

func ConsulChaeck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "consulCheck")
}

func RegistryServer() {
	config := consulapi.DefaultConfig()
	client, err := consulapi.NewClient(config)

	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	checkPort := 8080

	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "service-node-1"
	registration.Name = "service-node"
	registration.Port = 9527
	registration.Tags = []string{"service-node"}
	registration.Address = "127.0.0.1"
	registration.Check = &consulapi.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, checkPort, "/check"),
		Timeout:                        "3s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "30s",
	}

	err = client.Agent().ServiceRegister(registration)

	if err != nil {
		log.Fatal("registry server error : ", err)
	}

	http.HandleFunc("/check", ConsulChaeck)
	http.ListenAndServe(fmt.Sprintf(":%d", checkPort), nil)
}

func main() {
	RegistryServer()
}

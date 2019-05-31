package registry

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
)

func RegistryServer() {
	config := consulapi.DefaultConfig()
	client, err := consulapi.NewClient(config)

	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	checkPort := 8080

	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "account-service"
	registration.Name = "account-service"
	registration.Port = 8080
	registration.Tags = []string{"account-service"}
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
}

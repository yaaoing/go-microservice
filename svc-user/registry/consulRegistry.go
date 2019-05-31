package registry

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"time"
)

var log = logrus.New()

func RegistryServer() {
	config := consulapi.DefaultConfig()
	client, err := consulapi.NewClient(config)

	if err != nil {
		log.Error("consul client error : ", err)
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
		log.Error("registry server error : ", err)
	}

	go func() {
		for true {
			time.Sleep(10 * time.Second)
			_, _, err := client.Agent().AgentHealthServiceByID("account-service")
			if err != nil {
				log.Info("do registry again")
				if reconnect(client, registration) != nil {
					log.Fatal("cannot registry")
				}
			}
		}
	}()
}

func reconnect(client *consulapi.Client, registration *consulapi.AgentServiceRegistration) error {
	var err error
	for i := 0; i < 5; i++ {
		err = client.Agent().ServiceRegister(registration)

		if err != nil {
			log.Error("registry server error : ", err)
			time.Sleep(10 * time.Second)
		}
	}
	return err
}

package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
	"time"
)

var log = logrus.New()

var client = new(consulapi.Client)

func RegistryServer() {
	config := consulapi.DefaultConfig()
	var err error
	client, err = consulapi.NewClient(config)

	if err != nil {
		log.Error("consul client error : ", err)
	}

	checkPort := 8081

	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "product-service"
	registration.Name = "product-service"
	registration.Port = 8081
	registration.Tags = []string{"product-service"}
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

func GetClient() consulapi.Client {
	return *client
}

func main() {
	RegistryServer()
	client := GetClient()
	services, _, err := client.Health().Service("account-service", "", true, &consulapi.QueryOptions{})
	if err != nil {
		log.Warn("error retrieving instances from Consul: %v", err)
	}
	addrs := make(map[string]string)
	for _, service := range services {
		addrs[net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port))] = net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port))
	}

	for address := range addrs {
		log.Info(address)
	}
}

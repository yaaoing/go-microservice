package registry

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

type ZookeeperClient struct {
}

func (c ZookeeperClient) Registry() {
	hosts := []string{"localhost:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		log.Fatal(err)
	}

	serviceName := "/svc-account"
	category := "/provider"
	instance := []byte("localhost:8080")
	var flags int32 = 0
	acls := zk.WorldACL(zk.PermAll)

	isExist, _, err := conn.Exists(serviceName)
	if err != nil {
		log.Fatal(err)
	}
	if !isExist {
		data := []byte{}
		p, err := conn.Create(serviceName, data, flags, acls)
		if err != nil {
			log.Fatal(err)
		}
		log.Info("result path is: " + p)
	}
	p, err := conn.Create(serviceName+category, instance, zk.FlagEphemeral, acls)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("result path is: " + p)

	data, _, err := conn.Get(serviceName + category)
	if err != nil {
		log.Error(err)
	}
	log.Info("data: " + string(data))
}

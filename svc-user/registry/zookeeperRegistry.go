package registry

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

type ZookeeperClient struct {
}

func (c ZookeeperClient) RegistryServer() {
	hosts := []string{"localhost:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		log.Error(err)
	}
	defer conn.Close()

	serviceName := "/svc-user"
	category := "/provider"
	instance := []byte("localhost:8080")
	var flags int32 = 0
	acls := zk.WorldACL(zk.PermAll)

	p, err := conn.Create(serviceName+category, instance, flags, acls)
	if err != nil {
		log.Error(err)
	}
	log.Info("result path is: " + p)

}

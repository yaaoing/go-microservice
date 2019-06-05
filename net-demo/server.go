package main

import (
	"github.com/sirupsen/logrus"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:1024")
	if err != nil {
		logrus.Fatal("can not listen on 1024: ", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			logrus.Info("failed to accept")
			continue
		}
		logrus.Info("connected to client: ", conn.RemoteAddr())
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			logrus.Info("data read failed: ", err)
			return
		}
		logrus.Info("receiving data: ", string(buffer[:n]))
	}
}

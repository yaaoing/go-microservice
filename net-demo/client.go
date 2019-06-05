package main

import (
	"github.com/sirupsen/logrus"
	"net"
)

func main() {
	server := "127.0.0.1:1024"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		logrus.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		logrus.Fatal("failed to connect to server")
	}
	logrus.Info("connect to server successful")

	conn.Write([]byte("hello world!"))

}

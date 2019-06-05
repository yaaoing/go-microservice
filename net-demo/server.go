package main

import (
	"github.com/sirupsen/logrus"
	"leo/go-microservice/net-demo/protocol"
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
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	tmpBuffer := make([]byte, 0)

	//readerChannel := make(chan []byte, 16)
	//go reader(readerChannel)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			logrus.Info(err)
			return
		}
		tmpBuffer, err = protocol.Depack(append(tmpBuffer, buffer[:n]...))
		//logrus.Info(string(buffer[:n]))
		if err != nil {
			logrus.Fatal(err)
		}
		logrus.Info(string(tmpBuffer))
	}
	defer conn.Close()
}

func reader(readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			logrus.Info(string(data))
		}
	}
}

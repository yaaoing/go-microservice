package main

import (
	"github.com/sirupsen/logrus"
	"leo/go-microservice/net-demo/protocol"
	"net"
	"strconv"
	"time"
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

	send(conn)
}

func send(conn net.Conn) {
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		logrus.Fatal("con not convert net.Conn to net.TCPConn")
	}
	tcpConn.SetNoDelay(true)
	defer tcpConn.Close()
	for i := 0; i < 10; i++ {
		session := getSession()
		words := "{\"ID\":" + strconv.Itoa(i) + "\",\"Session\":" + session + "2015073109532345\",\"Meta\":\"golang\",\"Content\":\"message\"}"
		data, err := protocol.Enpack([]byte(words))
		if err != nil {
			logrus.Info(err)
		}
		conn.Write(data)
	}
}

func getSession() string {
	gs1 := time.Now().Unix()
	gs2 := strconv.FormatInt(gs1, 10)
	return gs2
}

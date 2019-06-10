package protocol

import (
	"bytes"
	"encoding/binary"
	"github.com/sirupsen/logrus"
)

const (
	Header        = "Herders"
	HeaderLength  = 7
	MessageLength = 4
)

func Enpack(message []byte) ([]byte, error) {
	b, err := IntToBytes(len(message))
	if err != nil {
		logrus.Info(err)
	}
	return append(append([]byte(Header), b...), message...), err
}

func Depack(buffer []byte) ([]byte, error) {
	length := len(buffer)

	var i int
	data := make([]byte, 32)
	for i = 0; i < length; i++ {
		// empty message
		if length < i+HeaderLength+MessageLength {
			break
		}
		// header message
		if string(buffer[i:i+HeaderLength]) == Header {
			msgLength, err := BytesToInt(buffer[i+HeaderLength : i+HeaderLength+MessageLength])
			if err != nil {
				return make([]byte, 0), err
			}
			if length < i+HeaderLength+MessageLength+msgLength {
				break
			}
			data = buffer[i+HeaderLength+MessageLength : i+HeaderLength+MessageLength+msgLength]
		}
	}
	if i == length {
		return make([]byte, 0), nil
	}
	return data, nil
}

func IntToBytes(n int) ([]byte, error) {
	x := int32(n)

	bytesBuf := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuf, binary.BigEndian, x)
	if err != nil {
		logrus.Info(err)
	}
	return bytesBuf.Bytes(), err
}

func BytesToInt(b []byte) (int, error) {
	bytesBuf := bytes.NewBuffer(b)

	var x int32
	err := binary.Read(bytesBuf, binary.BigEndian, &x)
	if err != nil {
		logrus.Info(err)
	}
	return int(x), err
}

package wlogn

import (
	"fmt"
	"net"
)

type Wlogn struct {
}

func New() *Wlogn {
	return &Wlogn{}
}

func (w *Wlogn) Listen(port int) error {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: port})
	if err != nil {
		return err
	}

	conn, err := listener.Accept()
	if err != nil {
		return err
	}
	var data []byte
	_, err = conn.Read(data)
	if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}

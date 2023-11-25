package wlogn

import (
	"bufio"
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
	reader := bufio.NewReader(conn)
	data, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}

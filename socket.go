package socket

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

import (
	"github.com/chuckpreslar/emission"
)

type Event uint8

const (
	Connection Event = iota
	Close
	Data
	Error
)

type Protocol string

const (
	TCP Protocol = "tcp"
	UDP Protocol = "udp"
)

const (
	PacketSize = 1024
)

type Socket struct {
	*emission.Emitter
	connection net.Conn
}

func (s *Socket) read() {
	for {
		data := make([]byte, PacketSize)

		if n, err := s.connection.Read(data); nil != err {
			if io.EOF == err {
				s.Emit(Close)
			} else {
				s.Emit(Error, err)
			}

			break
		} else if n > 0 {
			buffer := bytes.NewBuffer(data)
			s.Emit(Data, buffer)
		}
	}
}

func (s *Socket) On(event Event, listener interface{}) *Socket {
	s.Emitter.On(event, listener)
	return s
}

func (s *Socket) Write(bytes []byte) *Socket {
	go func(socket *Socket) {
		if _, err := socket.connection.Write(bytes); nil != err {
			socket.Emit(Error, err)
		}
	}(s)

	return s
}

func (s *Socket) WriteString(str string) *Socket {
	return s.Write([]byte(str))
}

func Connect(host string, port int, protocol Protocol, listener interface{}) {
	socket := new(Socket)
	socket.Emitter = emission.NewEmitter()
	socket.On(Connection, listener)

	var err error

	if socket.connection, err = net.Dial(string(protocol), fmt.Sprintf("%v:%v", host, port)); nil != err {
		socket.Emit(Connection, socket).Emit(Error, err)
	} else {
		socket.Emit(Connection, socket)
		socket.read()
	}
}

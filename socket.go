package socket

import (
	"bytes"
	"fmt"
	"net"
	"sync"
	"time"
)

import (
	"github.com/chuckpreslar/emission"
)

type event uint8

const (
	Connection event = iota
	Close
	Data
	Drain
	End
	Error
	Listening
	Timout
)

const (
	PacketSize = 1024
)

type Socket struct {
	*emission.Emitter
	*sync.WaitGroup
	net.Conn

	buffer    *bytes.Buffer
	connected bool
}

func (s *Socket) read() {
	for {
		data := make([]byte, PacketSize)
		s.SetDeadline(time.Now().Add(5 * time.Millisecond))

		if n, err := s.Read(data); nil != err {
			s.Emit(Error, err)
			break
		} else if n > 0 {
			buffer := bytes.NewBuffer(data)
			s.Emit(Data, buffer)
		} else {
			s.Emit(Drain)
		}
	}

	s.Done()
}

func (s *Socket) Connect(h string, p int, fn ...interface{}) *Socket {
	var err error

	if 0 < len(fn) {
		s.On(Connection, fn[0])
	}

	if s.Conn, err = net.Dial("tcp", fmt.Sprintf("%v:%v", h, p)); nil != err {
		s.Emit(Error, err)
	} else {
		s.Emit(Connection)
	}

	s.connected = true
	s.Add(1)
	defer s.Wait()

	go s.read()
	return s
}

func (s *Socket) On(e event, fn interface{}) *Socket {
	s.Emitter.On(e, fn)

	if e == Data {

	}

	return s
}

func (s *Socket) Close() *Socket { return s }

func NewSocket() (s *Socket) {
	s = new(Socket)
	s.WaitGroup = new(sync.WaitGroup)
	s.Emitter = emission.NewEmitter()
	s.buffer = &bytes.Buffer{}
	return
}

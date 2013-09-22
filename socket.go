package socket

import (
  "github.com/chuckpreslar/emission"
)

type Event uint8

const (
  CONNECTION Event = iota
  CLOSE
  DATA
  DRAIN
  END
  ERROR
  LISTENING
  TIMEOUT
)

type Listener interface {
  accept()
}

type EventListener func()
type ConnectionListener func(Socket)
type DataListener func(Buffer)

func (h EventListener) accept()      {}
func (h ConnectionListener) accept() {}
func (h DataListener) accept()       {}

type Socket struct {
  emitter *emission.Emitter
}

func (s *Server) Close() *Socket {
  return s
}

func NewSocket() (socket *Socket, err error) {
  socket = new(Socket)
  return
}

type Server struct {
  emitter *emission.Emitter
}

func (s *Server) Listen() *Server {
  return s
}

func (s *Server) Close() *Server {
  return s
}

func NewServer() (server *Server, err error) {
  server = new(Server)
  return
}

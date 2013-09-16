package socket

type Event uint8

const (
  CONNECT Event = iota
  CONNECTION
  CLOSE
  DATA
  DRAIN
  END
  ERROR
  LISTENING
  TIMEOUT
)

type Server struct{}
type Socket struct{}
type Buffer struct{}

type Handler interface {
  handler()
}

type EventHandler func() error
type ConnectionHandler func(Socket) error
type DataHandler func(Buffer) error

func (h EventHandler) handler()      {}
func (h ConnectionHandler) handler() {}
func (h DataHandler) handler()       {}

package socket

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

type Emitter struct {
  events map[Event][]Handler
}

func (e *Emitter) AddListener(event Event, handler Handler) {}
func (e *Emitter) On(event Event, handler Handler)          {}
func (e *Emitter) RemoveListener(handler Handler)           {}
func (e *Emitter) Off(handler Handler)                      {}

type Server struct {
  *Emitter
}

func (s *Server) Listen() (err error) {}
func (s *Server) Close() (err error)  {}

type Socket struct {
  *Emitter
}

package server

import (
	"os"
	"os/signal"
	"syscall"
)

type Handler interface {
	Start(stop <-chan struct{})
}

type Server struct {
	handlers []Handler
}

func New(opts ...Option) *Server {
	op := defaultOption()
	for _, f := range opts {
		f(op)
	}

	s := &Server{
		handlers: []Handler{},
	}
	s.handlers = append(s.handlers, newMCPServer(op))
	return s
}

func (s *Server) Start(stop <-chan struct{}) {
	for _, h := range s.handlers {
		go h.Start(stop)
	}
	<-stop
}

func SetupSignalHandler() <-chan struct{} {
	stop := make(chan struct{})

	c := make(chan os.Signal, 2)
	signal.Notify(c, syscall.SIGKILL, syscall.SIGTERM)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(-1)
	}()

	return stop
}

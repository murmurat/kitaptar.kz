package httpserver

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	server          *http.Server
	shutdownTimeout time.Duration
	notify          chan error
}

func New(handler http.Handler, opts ...Option) *Server {
	httpServer := &http.Server{Handler: handler}
	s := &Server{
		server: httpServer,
		notify: make(chan error, 1),
	}

	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)

	defer cancel()

	return s.server.Shutdown(ctx)
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

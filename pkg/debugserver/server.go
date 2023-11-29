package debugserver

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type BebugServer struct {
	server *http.Server
	notify chan error
}

func New(handler http.Handler, opts ...Option) *BebugServer {
	httpServer := &http.Server{Handler: handler}
	s := &BebugServer{
		server: httpServer,
		notify: make(chan error, 1),
	}

	for _, opt := range opts {
		opt(s)
	}

	fmt.Println(httpServer.Addr)
	return s
}

func Profiler() http.Handler {
	router := chi.NewRouter()

	router.Mount("/debug", middleware.Profiler())

	return router
}

func (s *BebugServer) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

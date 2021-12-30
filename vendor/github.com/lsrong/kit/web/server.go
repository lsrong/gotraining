package web

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"strings"
	"syscall"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type Server struct {
	shutdown chan os.Signal

	mux *httprouter.Router
	mws []Middleware
}

func NewServer(shutdown chan os.Signal, mws ...Middleware) *Server {
	return &Server{
		shutdown: shutdown,
		mux:      httprouter.New(),
		mws:      mws,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) SignalShutdown() {
	s.shutdown <- syscall.SIGTERM
}

func (s *Server) Use(mw ...Middleware) {
	s.mws = append(s.mws, mw...)
}

func (s *Server) Handle(method, group, path string, handler Handler, mws ...Middleware) {
	handler = wrapMiddleware(mws, handler)

	handler = wrapMiddleware(s.mws, handler)

	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if err := handler(ctx, w, r); err != nil {
			// need to shutdown
			s.SignalShutdown()
		}
	}
	fullPath := path
	if group != "" {
		fullPath = fmt.Sprintf("/%s/%s", strings.Trim(group, "/"), strings.TrimLeft(path, "/"))
	}

	s.mux.HandlerFunc(method, fullPath, h)
}

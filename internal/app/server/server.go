package server

import (
	"context"
	"net/http"
	"toky/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func New(cfg config.Server, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         ":" + cfg.Port,
			Handler:      handler,
			ReadTimeout:  cfg.Timeout,
			WriteTimeout: cfg.Timeout,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

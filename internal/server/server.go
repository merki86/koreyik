package server

import (
	"context"
	"github.com/serwennn/koreyik/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func New(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:        cfg.Server.Address,
			Handler:     handler,
			ReadTimeout: cfg.Server.Timeout,
			IdleTimeout: cfg.Server.IdleTimeout,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

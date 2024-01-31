package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"dev11/internal/config"
)

type Server struct {
	*http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(cfg config.Config, handler http.Handler) error {
	s.Server = &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      handler,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
	}
	return s.Server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}

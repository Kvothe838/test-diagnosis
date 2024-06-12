package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	_defaultHeaderTimeoutInSecs = 10
)

type Server struct {
	l   *logrus.Logger
	srv *http.Server
}

func New(port string) *Server {
	srv := &http.Server{
		Addr:              fmt.Sprintf("localhost:%s", port),
		ReadHeaderTimeout: time.Duration(_defaultHeaderTimeoutInSecs) * time.Second,
	}
	return &Server{
		srv: srv,
	}
}

func (s *Server) RegisterHandler(h http.Handler) {
	s.srv.Handler = h
}

func (s *Server) StartAsync() {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.l.Fatalf("server exiting: %s", err)
		}
	}()
}

func (s *Server) Close() error {
	return s.srv.Shutdown(context.Background())
}

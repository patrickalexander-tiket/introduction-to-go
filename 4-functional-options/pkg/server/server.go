package server

import (
	"log"
	"time"
)

type Server struct {
	Host    string
	Port    int
	Timeout time.Duration
	MaxConn int
}

type Option func(*Server)

func WithHost(host string) Option {
	return func(s *Server) {
		s.Host = host
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func WithMaxConn(maxConn int) Option {
	return func(s *Server) {
		s.MaxConn = maxConn
	}
}

func New(option ...Option) *Server {
	s := &Server{}
	for _, o := range option {
		o(s)
	}
	return s
}

func (s *Server) Start() {
	log.Printf("Server started %#v\n", s)
}

func (s *Server) Stop() {
	log.Printf("Server stopped %#v\n", s)
}

package gite

import (
	"net/http"
)

type Config struct {
	Logging bool
	Port    string
	OnServe func(port string)
}

func New(c *Config) *Server {
	mux := http.NewServeMux()
	return &Server{mux: mux, config: *c}
}

func (s *Server) Serve() error {
	s.config.OnServe(s.config.Port)
	e := http.ListenAndServe(join(":", s.config.Port), s.mux)
	return e
}

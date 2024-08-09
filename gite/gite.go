package gite

import (
	"net/http"
)

type Config struct {
	OnServe func(port string)
	Port    string
	Logging bool
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

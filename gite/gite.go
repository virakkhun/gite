package gite

import "net/http"

func NewServer(mux *http.ServeMux) *Server {
	return &Server{mux: mux}
}

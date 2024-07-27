package gite

import "net/http"

func New(mux *http.ServeMux) *Server {
	return &Server{mux: mux}
}

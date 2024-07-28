package gite

import (
	"net/http"
	"time"
)

type Server struct {
	mux    *http.ServeMux
	config Config
	ready  bool
}

type Router interface {
	Get(path string, handlers ...HanlderFunc)
	Post(path string, hanlders ...HanlderFunc)
	Put(path string, hanlders ...HanlderFunc)
	Patch(path string, hanlders ...HanlderFunc)
	Delete(path string, hanlders ...HanlderFunc)
	Group(prefix string) Router
}

func (s *Server) Get(path string, hanlders ...HanlderFunc) {
	p := &buildPathConfig{method: GET, path: path}
	s.registerHandlers(p.build(), hanlders...)
}

func (s *Server) Post(path string, handlers ...HanlderFunc) {
	p := &buildPathConfig{method: POST, path: path}
	s.registerHandlers(p.build(), handlers...)
}

func (s *Server) Put(path string, handlers ...HanlderFunc) {
	p := &buildPathConfig{method: PUT, path: path}
	s.registerHandlers(p.build(), handlers...)
}

func (s *Server) Patch(path string, handlers ...HanlderFunc) {
	p := &buildPathConfig{method: PATCH, path: path}
	s.registerHandlers(p.build(), handlers...)
}

func (s *Server) Delete(path string, handlers ...HanlderFunc) {
	p := &buildPathConfig{method: DELETE, path: path}
	s.registerHandlers(p.build(), handlers...)
}

func (s *Server) Group(prefix string) Router {
	g := &Group{app: s, prefix: prefix}
	return g
}

func (s *Server) toggleReady() {
	s.ready = true
}

func (s *Server) logger(r *http.Request) {
	logf("\n[%v] %v on %v", r.Method, r.URL.Path, time.Now().Format(time.UnixDate))
}

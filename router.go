package gite

import (
	"net/http"
)

type Server struct {
	mux    *http.ServeMux
	config Config
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
	p := &buildPathConfig{method: get, path: path}
	s.registerHandlers(p.build(), hanlders...)
}

func (s *Server) Post(path string, handlers ...HanlderFunc) {
	p := &buildPathConfig{method: post, path: path}
	s.registerHandlers(p.build(), handlers...)
}

func (s *Server) Put(path string, handlers ...HanlderFunc) {
	p := &buildPathConfig{method: put, path: path}
	s.registerHandlers(p.build(), handlers...)
}

func (s *Server) Patch(path string, handlers ...HanlderFunc) {
	p := &buildPathConfig{method: patch, path: path}
	s.registerHandlers(p.build(), handlers...)
}

func (s *Server) Delete(path string, handlers ...HanlderFunc) {
	p := &buildPathConfig{method: delete, path: path}
	s.registerHandlers(p.build(), handlers...)
}

func (s *Server) Group(prefix string) Router {
	g := &Group{app: s, prefix: prefix}
	return g
}

package gite

import (
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

type Router interface {
	Get(path string, handlers ...HanlderFunc)
	Post(path string, hanlders ...HanlderFunc)
	Put(path string, hanlders ...HanlderFunc)
	Patch(path string, hanlders ...HanlderFunc)
	Delete(path string, hanlders ...HanlderFunc)
	Group(prefix string) Router
}

func (s Server) Get(path string, hanlders ...HanlderFunc) {
	p := &buildPathConfig{method: GET, path: path}
	registerHandlers(s.mux, p.build(), hanlders...)
}

func (s Server) Post(path string, handlers ...HanlderFunc) {
	p := &buildPathConfig{method: POST, path: path}
	registerHandlers(s.mux, p.build(), handlers...)
}

func (s Server) Put(path string, handlers ...HanlderFunc) {
	p := &buildPathConfig{method: PUT, path: path}
	registerHandlers(s.mux, p.build(), handlers...)
}

func (s Server) Patch(path string, handlers ...HanlderFunc) {
	p := &buildPathConfig{method: PATCH, path: path}
	registerHandlers(s.mux, p.build(), handlers...)
}

func (s Server) Delete(path string, handlers ...HanlderFunc) {
	p := &buildPathConfig{method: DELETE, path: path}
	registerHandlers(s.mux, p.build(), handlers...)
}

func (s Server) Group(prefix string) Router {
	g := &Group{app: &s, prefix: prefix}
	return g
}

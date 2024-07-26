package gite

import (
	"net/http"
)

type Ctx struct {
	Response http.ResponseWriter
	Request  *http.Request
	NextFunc func()
}

type HanlderFunc func(ctx Ctx)

type Router struct {
	Get    func(path string, handlers ...HanlderFunc)
	Post   func(path string, handlers ...HanlderFunc)
	Patch  func(path string, handlers ...HanlderFunc)
	Put    func(path string, handlers ...HanlderFunc)
	Delete func(path string, handlers ...HanlderFunc)
	Group  func(prefix string) Router
}

package gite

import (
	"net/http"
)

type Router struct {
	Get    func(path string, handlers ...http.HandlerFunc)
	Post   func(path string, handlers ...http.HandlerFunc)
	Patch  func(path string, handlers ...http.HandlerFunc)
	Put    func(path string, handlers ...http.HandlerFunc)
	Delete func(path string, handlers ...http.HandlerFunc)
	Group  func(prefix string) Router
}

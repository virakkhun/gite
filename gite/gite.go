package gite

import "net/http"

func handler(mux *http.ServeMux, path string, handlers ...http.HandlerFunc) {
	for _, handler := range handlers {
		mux.Handle(path, handler)
	}
}

func GiteRouter(mux *http.ServeMux) Router {
	return Router{
		Get: func(path string, handlers ...http.HandlerFunc) {
			p := buildPath(HttpVerb.GET, path)
			handler(mux, p, handlers...)
		},
		Post: func(path string, handlers ...http.HandlerFunc) {
			p := buildPath(HttpVerb.POST, path)
			handler(mux, p, handlers...)
		},
		Put: func(path string, handlers ...http.HandlerFunc) {
			p := buildPath(HttpVerb.PUT, path)
			handler(mux, p, handlers...)
		},
		Patch: func(path string, handlers ...http.HandlerFunc) {
			p := buildPath(HttpVerb.PATCH, path)
			handler(mux, p, handlers...)
		},
		Delete: func(path string, handlers ...http.HandlerFunc) {
			p := buildPath(HttpVerb.DELETE, path)
			handler(mux, p, handlers...)
		},
		Group: func(prefix string) Router {
			router := GiteRouter(mux)
			return Router{
				Get: func(path string, handlers ...http.HandlerFunc) {
					p := buildGroupPath(prefix, path)
					router.Get(p, handlers...)
				},
				Post: func(path string, handlers ...http.HandlerFunc) {
					p := buildGroupPath(prefix, path)
					router.Post(p, handlers...)
				},
				Put: func(path string, handlers ...http.HandlerFunc) {
					p := buildGroupPath(prefix, path)
					router.Post(p, handlers...)
				},
				Patch: func(path string, handlers ...http.HandlerFunc) {
					p := buildGroupPath(prefix, path)
					router.Post(p, handlers...)
				},
				Delete: func(path string, handlers ...http.HandlerFunc) {
					p := buildGroupPath(prefix, path)
					router.Post(p, handlers...)
				},
				Group: func(prefix string) Router {
					router := GiteRouter(mux)
					return router.Group(prefix)
				},
			}
		},
	}
}

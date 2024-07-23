package gite

import "net/http"

func handler(mux *http.ServeMux, path string, handlers ...HanlderFunc) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		shuouldGoNext := false
		for _, handler := range handlers {
			handler(Ctx{
				Response: w,
				Request:  r,
				NextFunc: func() {
					shuouldGoNext = true
				},
			})

			if shuouldGoNext {
				continue
			} else {
				break
			}
		}
	})
}

func NewRouter(mux *http.ServeMux) Router {
	return Router{
		Get: func(path string, handlers ...HanlderFunc) {
			p := buildPath(HttpVerb.GET, path)
			handler(mux, p, handlers...)
		},
		Post: func(path string, handlers ...HanlderFunc) {
			p := buildPath(HttpVerb.POST, path)
			handler(mux, p, handlers...)
		},
		Put: func(path string, handlers ...HanlderFunc) {
			p := buildPath(HttpVerb.PUT, path)
			handler(mux, p, handlers...)
		},
		Patch: func(path string, handlers ...HanlderFunc) {
			p := buildPath(HttpVerb.PATCH, path)
			handler(mux, p, handlers...)
		},
		Delete: func(path string, handlers ...HanlderFunc) {
			p := buildPath(HttpVerb.DELETE, path)
			handler(mux, p, handlers...)
		},
		Group: func(prefix string) Router {
			router := NewRouter(mux)
			return Router{
				Get: func(path string, handlers ...HanlderFunc) {
					p := buildGroupPath(prefix, path)
					router.Get(p, handlers...)
				},
				Post: func(path string, handlers ...HanlderFunc) {
					p := buildGroupPath(prefix, path)
					router.Post(p, handlers...)
				},
				Put: func(path string, handlers ...HanlderFunc) {
					p := buildGroupPath(prefix, path)
					router.Post(p, handlers...)
				},
				Patch: func(path string, handlers ...HanlderFunc) {
					p := buildGroupPath(prefix, path)
					router.Post(p, handlers...)
				},
				Delete: func(path string, handlers ...HanlderFunc) {
					p := buildGroupPath(prefix, path)
					router.Post(p, handlers...)
				},
				Group: func(prefix string) Router {
					router := NewRouter(mux)
					return router.Group(prefix)
				},
			}
		},
	}
}

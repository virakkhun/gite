package gite

import (
	"net/http"
)

type Ctx struct {
	// raw response pass from HTTP hanlder
	//
	// @ref https://pkg.go.dev/net/http#ResponseWriter
	Response http.ResponseWriter
	// raw request pass from HTTP hanlder
	//
	// @ref https://pkg.go.dev/net/http#Request
	Request *http.Request
	// NewFunc
	//
	// a callback to trigger the next func handler
	NextFunc func()
}

type HanlderFunc func(ctx *Ctx)

func (s *Server) registerHandlers(path string, handlers ...HanlderFunc) {
	reportRegisteredHandlers(path, len(handlers))
	s.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if s.config.Logging {
			logger(r)
		}

		next := false

		for _, handler := range handlers {
			handler(&Ctx{
				Response: w,
				Request:  r,
				NextFunc: func() {
					next = true
				},
			})

			if next {
				continue
			} else {
				break
			}
		}
	})
}

package gite

import (
	"net/http"
)

type Ctx struct {
	Response http.ResponseWriter
	Request  *http.Request
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

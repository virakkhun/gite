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

func registerHandlers(mux *http.ServeMux, path string, handlers ...HanlderFunc) {
	reportRegisteredHandlers(path, len(handlers))
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

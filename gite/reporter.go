package gite

import (
	"fmt"
	"net/http"
	"time"
)

func reportRegisteredHandlers(path string, hanlders int) {
	fmt.Printf("%v => hanlders: %v\n", path, hanlders)
}

func logger(r *http.Request) {
	logf("\n[%v] %v on %v", r.Method, r.URL.Path, time.Now().Format(time.UnixDate))
}

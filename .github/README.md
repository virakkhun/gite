# gite

- is an minimal go router for create API.
- providing raw handler from net/http
- support route grouping
- support middleware

#### prerequisite

- go 1.22.5

#### basic routing

```go
package main

import (
	"encoding/json"
	"gite/gite"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	router := gite.NewRouter(mux)

	var myMiddleWare gite.HanlderFunc = func(ctx gite.Ctx) {
		if ctx.Request.URL.Path == "/hello" {
			http.NotFound(ctx.Response, ctx.Request)
			return
		}

		ctx.NextFunc()
	}

	router.Get("/", myMiddleWare, func(ctx gite.Ctx) {
		ctx.Response.WriteHeader(http.StatusCreated)
		fmt.Fprint(ctx.Response, "Hello world")
	})

	log.Fatal(http.ListenAndServe(":3000", mux))
}
```

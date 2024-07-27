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
	router := gite.New(mux)

	router.Get("/", func(ctx gite.Ctx) {
		json.NewEncoder(ctx.Response).Encode("Hello")
	})

	user := router.Group("/user")
	user.Get("/", func(ctx gite.Ctx) {
		if ctx.Request.URL.Path == "/user/detail" {
			http.NotFound(ctx.Response, ctx.Request)
			return
		}

		ctx.NextFunc()
	}, func(ctx gite.Ctx) {
		json.NewEncoder(ctx.Response).Encode("World")
	})

	log.Fatal(http.ListenAndServe(":3000", mux))
}
```

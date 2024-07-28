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
	app := gite.New(&gite.Config{
		Port:    "3000",
		Logging: true,
		OnServe: func(port string) {
			fmt.Printf("Server is listening on port %v\n", port)
		},
	})

	app.Get("/", func(ctx gite.Ctx) {
		ctx.Json("Hello World")
	})

	user := app.Group("/user")
	user.Get("/", func(ctx gite.Ctx) {
		if ctx.Request.URL.Path == "/user/detail" {
			http.NotFound(ctx.Response, ctx.Request)
			return
		}

		ctx.NextFunc()
	}, func(ctx gite.Ctx) {
		ctx.Status(http.StatusOK).Text("Hello bro")
	})

	log.Fatal(app.Serve())
}
```

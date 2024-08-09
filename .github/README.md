# gite

- is an minimal go router for create API.
- providing raw handler from net/http
- support route grouping
- support middleware

#### getting started

to get started, you will need `go 1.22.5` or later.
visit [go download page](https://go.dev/dl/)

init your project by,
`go mod init github.com/your/repo`

and then get `gite` lib
`go get github.com/virakkhun/gite@v0.0.1-beta`

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
  // config the app by passing the struct
	app := gite.New(&gite.Config{
		Port:    "3000",
		Logging: true,
		OnServe: func(port string) {
			fmt.Printf("Server is listening on port %v\n", port)
		},
	})

  // route to localhost:3000
  // get json value back
  // GET /
	app.Get("/", func(ctx gite.Ctx) {
		ctx.Json("Hello World")
	})

  // grouping for user
	user := app.Group("/user")

  // middleware
  testMiddleware := func(ctx gite.Ctx) {
		if ctx.Request.URL.Path == "/user/detail" {
			http.NotFound(ctx.Response, ctx.Request)
			return
		}

		ctx.NextFunc()
	}

  // route to /user/
  // GET /user/
	user.Get("/", testMiddleware, func(ctx gite.Ctx) {
		ctx.Status(http.StatusOK).Text("Hello bro")
	})

  // serve the application
	log.Fatal(app.Serve())
}
```

---

# 🧾 License

MIT License

Copyright (c) 2024 Virak Khun

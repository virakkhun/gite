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
  "github.com/virakkhun/gite"
  "log"
  "http"
)

func main() {
  // config the app by passing the struct
  app := gite.New(&gite.Config{
    Port: "3000",
    Logging: true
  })

  // GET /hello
  app.Get("/hello", func(ctx *gite.Ctx) {
    ctx.Json("Hello, world!!")
  })

  // create grouping
  user := app.Group("/user")

  // GET /user/{id}
  user.Get("/{id}", func(ctx *gite.Ctx) {
    ctx.Json(ctx.FromPath("id"))
  })

  // middleware
  // need to return and call NextFunc for middleware
  middleware := func(ctx *gite.Ctx) {
    if ctx.FromPath("id") == -1 {
      http.NotFound(ctx.Response, ctx.Request)
      return
    }

    ctx.NextFunc()
  }

  // GET /user/{id}
  // add middleware
  user.Get("/{id}", middleware, func(ctx *gite.Ctx) {
		ctx.Status(http.StatusOK).Text("Hello bro")
  })

  // serve the application
  log.Fatal(app.Serve())
}
```

# 🧾 License

MIT License

Copyright (c) 2024 Virak Khun

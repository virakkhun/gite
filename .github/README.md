# gite

is an minimal router for create API. allowing for group routing, and others http verbs.

```go
package main

import (
	"encoding/json"
	"gite/gite"
	"log"
	"net/http"
)

func main() {
	PORT := ":3000"
	mux := http.NewServeMux()

	router := gite.GiteRouter(mux)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(
			map[string]string{
				"Hello": "World",
			},
		)
	})

	// grouping
	user := router.Group("/user")

	user.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(
			map[string]any{
				"username": "@johndoe",
				"Age":      20,
			},
		)
	})

	err := http.ListenAndServe(PORT, mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}
```
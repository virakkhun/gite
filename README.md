# gite

is an minimal router for create API.
allowing for group routing, and others http verbs.

```ts
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
  router.Post("/", handler)
  router.Put("/{id}", handler)
  router.Patch("/{id}", handler)
  router.Delete("/{id}", handler)

	err := http.ListenAndServe(PORT, mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// grouping
router := gite.GiteRouter(mux)
user := router.Group("/user")

user.Get("/", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(
		map[string]any{
			"username": "@johndoe",
      "Age": 20,
		},
	)
})
```

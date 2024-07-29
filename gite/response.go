package gite

import (
	"encoding/json"
	"fmt"
)

type Chainable interface {
	Json(data interface{})
	Text(data string)
	Status(statusCode int) Chainable
}

func (r Ctx) Json(data interface{}) {
	r.AddHeader(CONTENT_TYPE, APPLICATION_JSON)
	json.NewEncoder(r.Response).Encode(data)
}

func (r Ctx) Text(data string) {
	fmt.Fprint(r.Response, data)
}

func (r Ctx) AddHeader(key string, value string) {
	r.Response.Header().Add(key, value)
}

func (r Ctx) DeleteHeader(key string) {
	r.Response.Header().Del(key)
}

func (r Ctx) GetHeader(key string) string {
	return r.Response.Header().Get(key)
}

func (r Ctx) Status(statusCode int) Chainable {
	r.Response.WriteHeader(statusCode)
	return &r
}

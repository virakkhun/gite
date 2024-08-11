package gite

import (
	"encoding/json"
	"io"
	"net/http"
)

// Response interface contains Json, Text and Status method.
// implemented for struct Ctx
type Response interface {
	Json(data interface{})
	Text(data string)
	Status(statusCode int) Response
}

// response json data to the consumer
//
// example:
//
//	ctx.Json(map[string]string{
//	  "Hello": "World"
//	})
func (r *Ctx) Json(data interface{}) {
	r.AddHeader(content_type, application_json)
	json.NewEncoder(r.Response).Encode(data)
}

// response text string to the consumer
//
// example:
//
// ctx.Text("Hello")
func (r *Ctx) Text(data string) {
	io.WriteString(r.Response, data)
}

// add header to the current response context
//
// example:
//
// ctx.AddHeader("my-header", "my value")
func (r *Ctx) AddHeader(key string, value string) {
	r.Response.Header().Add(key, value)
}

// delete a header from the current context
//
// example:
//
// ctx.DeleteHeader("my-header")
func (r *Ctx) DeleteHeader(key string) {
	r.Response.Header().Del(key)
}

// get a specific header
//
// example:
//
// ctx.GetHeader("my-header")
func (r *Ctx) GetHeader(key string) string {
	return r.Response.Header().Get(key)
}

// Status return a Response interface
// which allow adding status to the current context
// and resonse data to the consumer by calling Json or Text
//
// example:
//
// - ctx.Status(http.StatusCreated)
//
// - ctx.Status(http.StatusOk).Json("Hello world!!")
func (r *Ctx) Status(statusCode int) Response {
	r.Response.WriteHeader(statusCode)
	return r
}

func (r *Ctx) ThrowNotFound() {
	http.NotFound(r.Response, r.Request)
}

func (r *Ctx) ThrowBadRequest() {
	r.Status(http.StatusBadRequest).Text("Bad request")
}

func (r *Ctx) ThrowUnauthorized() {
	r.Status(http.StatusUnauthorized).Text("Unauthorized")
}

func (r *Ctx) ThrowUnprocessableEntity() {
	r.Status(http.StatusUnprocessableEntity).Text("Unprocessable entity")
}

func (r *Ctx) ThrowInternalServerError() {
	r.Status(http.StatusInternalServerError).Text("Internal server error")
}

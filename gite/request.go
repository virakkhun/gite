package gite

import (
	"encoding/json"
)

// ToJson
// parse request body to json
//
// example:
//
//		data := new(map[string]string)
//		err := ctx.GetBodyAsJson(&data)
//
//	---
//
//		type Person struct {
//		 name string
//		 age number
//		}
//		person := new(&Person{})
//		err := ctx.GetBodyAsJson(&person)
func (c *Ctx) ToJson(data interface{}) error {
	err := json.NewDecoder(c.Request.Body).Decode(data)
	return err
}

// FromPath
// get value from path
// return empty string if it doesn't match
//
// @ref https://pkg.go.dev/net/http#Request.PathValue
//
// example:
//
//	app.Get("/users/{userId}", func(c *Ctx) {
//	   userId := c.FromPath("userId")
//	})
func (c *Ctx) FromPath(name string) string {
	return c.Request.PathValue(name)
}

// FromQuery
// get value form query
// return emtry string if it doesn't match
//
// @ref https://pkg.go.dev/net/url#Values.Get
//
// example:
//
//	app.Get("/users", func(c *Ctx) {
//	  currency := c.FromQuery("currency")
//	})
func (c *Ctx) FromQuery(key string) string {
	q := c.Request.URL.Query()
	return q.Get(key)
}

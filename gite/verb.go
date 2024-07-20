package gite

type verb struct {
	GET    string
	POST   string
	PUT    string
	PATCH  string
	DELETE string
}

var HttpVerb verb = verb{
	GET:    GET,
	POST:   POST,
	PATCH:  PATCH,
	PUT:    PUT,
	DELETE: DELETE,
}

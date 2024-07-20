package gite

import (
	"strings"
	"testing"
)

func TestBuildPath(t *testing.T) {
	get := buildPath(HttpVerb.GET, "/")

	if strings.Compare(get, "GET /") != 0 {
		t.Fatal(`failed to build GET /`)
	}
	t.Log(`Results: `, get)

	post := buildPath(HttpVerb.POST, "/")
	if strings.Compare(post, "POST /") != 0 {
		t.Fatal(`failed to build POST /`)
	}
	t.Log(`Results: `, post)

	put := buildPath(HttpVerb.PUT, "/")
	if strings.Compare(put, "PUT /") != 0 {
		t.Fatal(`failed to build PUT /`)
	}
	t.Log(`Results: `, put)

	patch := buildPath(HttpVerb.PATCH, "/")
	if strings.Compare(patch, "PATCH /") != 0 {
		t.Fatal(`failed to build PATCH /`)
	}
	t.Log(`Results: `, patch)

	del := buildPath(HttpVerb.DELETE, "/")
	if strings.Compare(del, "DELETE /") != 0 {
		t.Fatal(`failed to build DELETE /`)
	}
	t.Log(`Results: `, del)
}

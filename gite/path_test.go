package gite

import (
	"testing"
)

func TestBuildPath(t *testing.T) {
	get := &buildPathConfig{method: GET}
	result := get.build()

	if result != "GET /" {
		t.Fatal(`failed to build GET /`)
	}
	t.Logf(`build %v success`, result)

	post := &buildPathConfig{method: POST}
	postResult := post.build()
	if postResult != "POST /" {
		t.Fatal(`failed to build POST /`)
	}
	t.Logf(`build %v success`, postResult)

	patch := &buildPathConfig{method: PATCH, path: "/{userId}"}
	patchResult := patch.build()
	if patchResult != "PATCH /{userId}" {
		t.Fatal(`failed to build PATCH /`)
	}
	t.Logf(`build %v success`, patchResult)
}

func TestBuildGroupPath(t *testing.T) {
	get := &buildPathConfig{prefix: "/user", path: "/"}
	expected := "/user/"
	result := get.buildGroup()

	if result != expected {
		t.Fatalf("failed to build group GET /user/")
	}

	t.Logf(`build group %v success`, result)

	post := &buildPathConfig{prefix: "/user"}
	postExpected := "/user/"
	postResult := post.buildGroup()

	if postExpected != postResult {
		t.Fatalf("failed to build group GET /user/")
	}

	t.Logf(`build group %v success`, postResult)

	put := &buildPathConfig{prefix: "/user", path: "/{id}"}
	putExpected := "/user/{id}"
	putResult := put.buildGroup()

	if putExpected != putResult {
		t.Fatalf("failed to build group GET /user/")
	}

	t.Logf(`build group %v success`, putResult)
}

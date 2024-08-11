package gite

import "strings"

const seperator = " "

type buildPathConfig struct {
	method string
	prefix string
	path   string
}

func buildPath(method string, path string) string {
	return strings.Join([]string{method, seperator, path}, "")
}

func buildGroupPath(c buildPathConfig) string {
	return strings.Join([]string{c.method, seperator, c.prefix, c.path}, "")
}

func join(s ...string) string {
	return strings.Join(s, "")
}

func (b *buildPathConfig) build() string {
	if b.path == "" {
		return join(b.method, seperator, "/")
	}
	return join(b.method, seperator, b.path)
}

func (b *buildPathConfig) buildGroup() string {
	if b.prefix == "" {
		panic("prefix is empty")
	}

	if b.path == "" {
		return join(b.prefix, "/")
	}

	return join(b.prefix, b.path)
}

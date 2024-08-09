package gite

import "strings"

const SEPERATOR = " "

type buildPathConfig struct {
	method string
	prefix string
	path   string
}

func buildPath(method string, path string) string {
	return strings.Join([]string{method, SEPERATOR, path}, "")
}

func buildGroupPath(c buildPathConfig) string {
	return strings.Join([]string{c.method, SEPERATOR, c.prefix, c.path}, "")
}

func join(s ...string) string {
	return strings.Join(s, "")
}

func (b *buildPathConfig) build() string {
	if b.path == "" {
		return join(b.method, SEPERATOR, "/")
	}
	return join(b.method, SEPERATOR, b.path)
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

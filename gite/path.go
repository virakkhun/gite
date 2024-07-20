package gite

import "strings"

const SEPERATOR = " "

func buildPath(method string, path string) string {
	return strings.Join([]string{method, SEPERATOR, path}, "")
}

func buildGroupPath(prefix string, path string) string {
	return strings.Join([]string{prefix, path}, "")
}

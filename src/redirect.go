package main

import (
	"net/http"
	"strings"
)

// 重定向
// file:redirect.go

var old = "/setu/latest"
var new = "/setu/v" + conf.Info.Version

func redirect(r *http.Request) bool {
	if r.URL.Path[:len(old)] == old && (len(old) == len(r.URL.Path) || []rune(r.URL.Path)[len(old)] == '/') {
		r.URL.Path = strings.Replace(r.URL.Path, old, new, 1)
		return true
	}
	if r.URL.Path[:len(new)] == new && (len(new) == len(r.URL.Path) || []rune(r.URL.Path)[len(old)] == '/') {
		return true
	}
	return false
}

package main

import (
	"strings"
)

// 重定向
// file:redirect.go

var old = "/setu/latest"
var new = "/setu/v" + config.Version

func redirect(url string) string {
	if strings.Index(url, old) == 0 {
		return strings.Replace(url, old, new, 1)
	}
	if strings.Index(url, new) == 0 {
		return url
	}
	return ""
}

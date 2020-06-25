package main

import (
	"strings"
)

// 重定向
// file:redirect.go

var oldUri = "/setu/latest"
var newUri = "/setu/v" + config.Version

func redirect(url string) string {
	if strings.Index(url, oldUri) == 0 {
		return strings.Replace(url, oldUri, newUri, 1)
	}
	if strings.Index(url, newUri) == 0 {
		return url
	}
	return ""
}

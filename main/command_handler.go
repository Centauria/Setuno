package main

import (
	"fmt"
	"net/http"
	"strings"
)

// 判断处理命令
// file:command_handler.go

type command struct {
	url     string
	method  string
	handler func(*http.Request, http.ResponseWriter) bool
}

var command_list = []command{
	command{"/info", "GET", info_handler},
	command{"/view", "GET", view_handler},
	command{"/upload", "POST", upload_handler},
}

func command_judge(r *http.Request, w http.ResponseWriter) bool {
	url := r.URL.Path[len(new):]
	method := r.Method
	for _, c := range command_list {
		if method == c.method && strings.Index(url, c.url) == 0 {
			if !c.handler(r, w) {
				return false
			}
			return true
		}
	}
	//404：找不到资源地址
	fmt.Println("StatusCode:404, Can't find command \"" + url + "\"")
	w.WriteHeader(404)
	return false
}

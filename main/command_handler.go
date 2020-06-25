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
	handler func(string, http.ResponseWriter) bool
}

var commandList = []command{
	{"/info", "GET", infoHandler},
	{"/view", "GET", viewHandler},
	{"/upload", "POST", uploadHandler},
}

func commandJudge(url string, method string, w http.ResponseWriter) bool {
	for _, c := range commandList {
		if method == c.method && strings.Index(url, c.url) == 0 {
			if !c.handler(url, w) {
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

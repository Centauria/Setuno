package main

import (
	"fmt"
	"net/http"
)

//按照命令处理 view
//file:view_handler.go

//TODO: 失败返回状态码
func viewHandler(url string, w http.ResponseWriter) bool {
	fmt.Println(url)
	return true
}

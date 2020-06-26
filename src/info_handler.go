package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//按照命令处理 info
func info_handler(r *http.Request, w http.ResponseWriter) bool {

	url := r.URL.String()[len(new):]

	if url == "/info" {
		msg, _ := json.Marshal(config)
		fmt.Println("StatusCode:200, Command \"" + url + "\", Server's information responded")
		_, _ = w.Write(msg)
		return true
	}

	//404：找不到资源地址
	fmt.Println("StatusCode:404, Can't find command \"" + url[5:] + "\" in info")
	w.WriteHeader(404)
	return false

}

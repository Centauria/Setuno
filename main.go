package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

// 重定向
// file:redirect.go

var old = "/setu/latest"
var new = "/setu/v0.1" + config.Version

func redirect(url string) string {
	if strings.Index(url, old) == 0 {
		return strings.Replace(url, old, new, 1)
	}
	if strings.Index(url, new) == 0 {
		return url
	}
	return ""
}

// 判断处理命令
// file:command_handler.go

type command struct {
	url     string
	method  string
	handler func(string, http.ResponseWriter) bool
}

var command_list = []command{
	command{"/info", "GET", info_handler},
	command{"/view", "GET", view_handler},
	command{"/upload", "POST", upload_handler},
}

func command_judge(url string, method string, w http.ResponseWriter) bool {
	for _, c := range command_list {
		if method == c.method && strings.Index(url, c.url) == 0 {
			if !c.handler(url, w) {
				return false
			}
			return true
		}
	}
	//404：找不到资源地址
	w.WriteHeader(404)
	return false
}

// 配置文件
// file: config.go
type config_list struct {
	Go_version string `json:"go_version"`
	Os         string `json:"os"`
	Arch       string `json:"arch"`
	Version    string `json:"version"`
}

var config = config_list{runtime.Version()[2:], runtime.GOOS, runtime.GOARCH, "0.1"}

//按照命令处理 info
//file:info_handler.go

//TODO: 失败返回状态码
func info_handler(url string, w http.ResponseWriter) bool {

	if url == "/info" {
		msg, _ := json.Marshal(config)
		w.Write(msg)
		return true
	}

	//404：找不到资源地址
	w.WriteHeader(404)
	return false

}

//按照命令处理 view
//file:view_handler.go

//TODO: 失败返回状态码
func view_handler(url string, w http.ResponseWriter) bool {
	fmt.Println(url)
	return true
}

//按照命令处理 upload
//file:upload_handler.go

//TODO: 失败返回状态码
func upload_handler(url string, w http.ResponseWriter) bool {
	return true
}

func setu_handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	//重定向判定
	url := redirect(r.URL.Path)
	if url == "" {
		//404:找不到资源地址
		w.WriteHeader(404)
		return
	}

	//去除头
	path := url[len(new):]

	//命令判断与处理
	if !command_judge(path, r.Method, w) {
		return
	}

	//w.Write([]byte("请求成功!!!\n"))
}

func main() {

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", setu_handler)

	err := http.ListenAndServe("127.0.0.1:9000", serveMux)
	if err != nil {
		fmt.Printf("http.ListenAndServe()函数执行错误,错误为:%v\n", err)
		return
	}
}

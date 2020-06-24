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

var oldUri = "/setu/latest"
var newUri = "/setu/v0.1" + config.Version

func redirect(url string) string {
	if strings.Index(url, oldUri) == 0 {
		return strings.Replace(url, oldUri, newUri, 1)
	}
	if strings.Index(url, newUri) == 0 {
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
	w.WriteHeader(404)
	return false
}

// 配置文件
// file: config.go
type configList struct {
	GoVersion string `json:"go_version"`
	Os        string `json:"os"`
	Arch      string `json:"arch"`
	Version   string `json:"version"`
}

var config = configList{runtime.Version()[2:], runtime.GOOS, runtime.GOARCH, "0.1"}

//按照命令处理 info
//file:info_handler.go

//TODO: 失败返回状态码
func infoHandler(url string, w http.ResponseWriter) bool {

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
func viewHandler(url string, w http.ResponseWriter) bool {
	fmt.Println(url)
	return true
}

//按照命令处理 upload
//file:upload_handler.go

//TODO: 失败返回状态码
func uploadHandler(url string, w http.ResponseWriter) bool {
	return true
}

func setuHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	//重定向判定
	url := redirect(r.URL.Path)
	if url == "" {
		//404:找不到资源地址
		w.WriteHeader(404)
		return
	}

	//去除头
	path := url[len(newUri):]

	//命令判断与处理
	if !commandJudge(path, r.Method, w) {
		return
	}

	//w.Write([]byte("请求成功!!!\n"))
}

func main() {

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", setuHandler)

	err := http.ListenAndServe("127.0.0.1:9000", serveMux)
	if err != nil {
		fmt.Printf("http.ListenAndServe()函数执行错误,错误为:%v\n", err)
		return
	}
}

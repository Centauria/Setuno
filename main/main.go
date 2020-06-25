package main

import (
	"fmt"
	"net/http"
)

func setuHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	//重定向判定
	url := redirect(r.URL.Path)
	if url == "" {
		//404:找不到资源地址
		fmt.Println("StatusCode:404, Can't find command \"" + r.URL.Path + "\"")
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

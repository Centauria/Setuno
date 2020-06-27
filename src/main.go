package main

import (
	"fmt"
	"net/http"
)

func setuHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	//重定向判定

	if !redirect(r) {
		//404:找不到资源地址
		fmt.Println(r.URL.String())
		fmt.Println("StatusCode:404, Can't find command \"" + r.URL.Path + "\"")
		w.WriteHeader(404)
		return
	}

	//命令判断与处理
	if !commandJudge(r, w) {
		return
	}

}

func main() {

	initConf()

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", setuHandler)

	err := http.ListenAndServe("127.0.0.1:9000", serveMux)
	if err != nil {
		fmt.Printf("http.ListenAndServe()函数执行错误,错误为:%v\n", err)
		return
	}
}

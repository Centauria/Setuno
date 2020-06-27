package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//按照命令处理 view
func viewHandler(r *http.Request, w http.ResponseWriter) bool {

	url := r.URL.String()[len(new)+len("/view"):]

	if r.URL.String()[len(new):] == "/view" || []rune(url)[0] == '?' {

		// 命令 /view 的处理
		return viewHandlerQuery(url, r, w)

	}

	if []rune(url)[0] == '/' {

		com := strings.Split(url[1:], "/")

		// 命令 /view/direct 的处理
		if len(com) == 2 && com[0] == "direct" {
			return viewHandlerDirect(url, com[1], w)
		}

		// 命令 /view/status 的处理
		if len(com) == 2 && com[0] == "status" {
			return viewHandlerStatus(url, com[1], w)
		}

		// 命令 /view/random 的处理
		bool1 := len(com) == 1 && (com[0] == "random" || (len(com[0]) > 6 && []rune(com[0])[6] == '?'))
		bool2 := len(com) == 2 && com[1] == ""
		if bool1 || bool2 {
			return viewHandlerRandom(url, r, w)
		}

	}

	//404：找不到资源地址
	fmt.Println("StatusCode:404, Can't find command " + url)
	w.WriteHeader(404)
	return false
}

// 命令/view? 的处理
func viewHandlerQuery(url string, r *http.Request, w http.ResponseWriter) bool {

	//查询
	query := r.URL.Query()
	ids, err := getIdByGet(query)
	if err != nil {
		fmt.Println("StatusCode:404, ", err)
		w.WriteHeader(404)
		return false
	}

	//Json化
	idsJson, err := json.Marshal(ids)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	_, _ = w.Write(idsJson)
	fmt.Println("StatusCode:200, Command \"" + url + "\", Server's information responded")

	return true
}

// 命令 /view/direct 的处理
func viewHandlerDirect(url string, id string, w http.ResponseWriter) bool {
	//获得路径
	path, err := getImageById(id)
	if err != nil {
		fmt.Println("StatusCode:404, ", err)
		w.WriteHeader(404)
		return false
	}

	//发送图片
	err = sendImage(path, w)
	if err != nil {
		fmt.Println("StatusCode:404, ", err)
		w.WriteHeader(404)
		return false
	}

	fmt.Println("StatusCode:200, Command \"" + url + "\", Server's information responded")
	return true
}

// 命令 /view/status 的处理
func viewHandlerStatus(url string, id string, w http.ResponseWriter) bool {
	statusJson, err := getImageStatusById(id)
	if err != nil {
		fmt.Println("StatusCode:404, ", err)
		w.WriteHeader(404)
		return false
	}
	_, _ = w.Write(statusJson)
	fmt.Println("StatusCode:200, Command \"" + url + "\", Server's information responded")
	return true
}

// 命令 /view/random 的处理
func viewHandlerRandom(url string, r *http.Request, w http.ResponseWriter) bool {
	query := r.URL.Query()

	//获取query为字符串
	qType := ""
	if query["type"] != nil {
		qTypeMap, _ := query["type"]
		qType = qTypeMap[0]
	}

	//获取随即图片路径
	path, err := getImageRandom(qType)
	if err != nil {
		fmt.Println("StatusCode:404, ", err)
		w.WriteHeader(404)
		return false
	}

	//发送图片
	err = sendImage(path, w)
	if err != nil {
		fmt.Println("StatusCode:404, ", err)
		w.WriteHeader(404)
		return false
	}

	fmt.Println("StatusCode:200, Command \"" + url + "\", Server's information responded")
	return true
}

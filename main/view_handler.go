package main

import (
	"fmt"
	"net/http"
)

/*
/view?range=$range&sort=$sort --> JSON
返回所有可查看的图片ID，以JSON列表的格式
$range
'' --> All pictures
':10' --> Pic [0, 1, ..., 9]
'10:' --> Pic [10, 11, ..., end]
'10:20' --> Pic [10, 11, ..., 19]
$sort
'', 'D' --> 时间新的在前
'A' --> 时间旧的在前
*/

//按照命令处理 view
//file:view_handler.go

//TODO: 失败返回状态码
func view_handler(r *http.Request, w http.ResponseWriter) bool {
	url := r.URL.String()[len(new)+len("/view"):]
	fmt.Println(url)
	if []rune(url)[0] == '/' {
		fmt.Println(url)
	}
	if []rune(url)[0] == '?' {
		fmt.Println(r.URL.Query())
	}
	return true
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

	if []rune(url)[0] == '/' {

		path, err := getImageById(url[1:])
		if err != nil {
			fmt.Println("StatusCode:404, ", err)
			w.WriteHeader(404)
			return false
		}

		err = sendImage(path, w)
		if err != nil {
			fmt.Println("StatusCode:404, ", err)
			w.WriteHeader(404)
			return false
		}
	}

	if []rune(url)[0] == '?' {
		fmt.Println(r.URL.Query())
	}
	return true
}

//根据_id查询单张图片
func getImageById(id string) (string, error) {

	//连接
	client, err := connectMongo()
	if err != nil {
		return "", err
	}

	// 指定获取要操作的数据集
	collectionLink := "setu_image"
	collection := client.Database(mongodb_link.db).Collection(collectionLink)
	fmt.Println("Connected to " + mongodb_link.db + "!")

	//按照_id查询图片地址
	result, err := findById(id, collection)
	if err != nil {
		return "", err
	}
	path := getUrlByResult(result)

	//断开连接
	err = disconnectMongo(client)
	if err != nil {
		return "", err
	}

	return path, nil
}

//传输图片
func sendImage(path string, w http.ResponseWriter) error {

	file, _ := os.Open(path)
	defer file.Close()
	buff, _ := ioutil.ReadAll(file)
	w.Write(buff)

	fmt.Println(buff)

	return nil
}

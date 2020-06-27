package main

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type imageNum struct {
	Number int `json:"number"`
}

//按照命令处理 count
func countHandler(r *http.Request, w http.ResponseWriter) bool {

	url := r.URL.String()[len(new)+len("/count"):]

	if r.URL.String()[len(new):] == "/count" || []rune(url)[0] == '?' {

		return countHandlerQuery(url, r, w)

	}

	//404：找不到资源地址
	fmt.Println("StatusCode:404, Can't find command \"" + url[5:] + "\" in info")
	w.WriteHeader(404)
	return false
}

// 命令/count? 的处理
func countHandlerQuery(url string, r *http.Request, w http.ResponseWriter) bool {

	//查询
	query := r.URL.Query()

	//获取query为字符串
	qType := ""
	if query["type"] != nil {
		qTypeMap, _ := query["type"]
		qType = qTypeMap[0]
	}

	//连接
	client, err := connectMongo()
	if err != nil {
		fmt.Println("StatusCode:404, ", err)
		w.WriteHeader(404)
		return false
	}

	// 指定获取要操作的数据集
	collectionLink := conf.Mongodb.Collection
	collection := client.Database(conf.Mongodb.Db).Collection(collectionLink)
	fmt.Println("Connected to " + conf.Mongodb.Db + "!")

	//获取图库图片数量
	var imageNumber imageNum
	if qType == "" {
		imageNumber.Number = countNum(collection, bson.D{{}})
	} else {
		imageNumber.Number = countNum(collection, bson.D{{"info.content", qType}})
	}

	//断开连接
	err = disconnectMongo(client)
	if err != nil {
		fmt.Println("StatusCode:404, ", err)
		w.WriteHeader(404)
		return false
	}

	//转为json
	msg, _ := json.Marshal(imageNumber)
	fmt.Println("StatusCode:200, Command \"" + url + "\", Server's information responded")
	_, _ = w.Write(msg)

	return true

}

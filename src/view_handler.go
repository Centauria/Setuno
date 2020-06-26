package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

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

		return true
	}

	if []rune(url)[0] == '?' {

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
		fmt.Println("StatusCode:200, Command \"" + url + "\", Server's information responded")
		_, _ = w.Write(idsJson)

		return true
	}

	//404：找不到资源地址
	fmt.Println("StatusCode:404, Can't find command \"" + url[5:] + "\" in info")
	w.WriteHeader(404)
	return false
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

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	buff, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	_, err = w.Write(buff)
	if err != nil {
		return err
	}

	return nil
}

//根据get参数获取_id
func getIdByGet(query map[string][]string) ([]string, error) {

	//连接
	client, err := connectMongo()
	if err != nil {
		return nil, err
	}

	// 指定获取要操作的数据集
	collectionLink := "setu_image"
	collection := client.Database(mongodb_link.db).Collection(collectionLink)
	fmt.Println("Connected to " + mongodb_link.db + "!")

	//得到正确qyery
	indexMin, indexMax, qSort, err := judgeAndFormatQuert(query, collection)
	if err != nil {
		return nil, err
	}

	//查找_id

	var ids []string
	if qSort == "A" {
		ids, err = findIdByRangeA(indexMin, indexMax, collection)
		if err != nil {
			return ids, err
		}
	}
	if qSort == "" || qSort == "D" {
		ids, err = findIdByRangeD(indexMin, indexMax, collection)
		if err != nil {
			return ids, err
		}
	}

	return ids, nil
}

//获得正确query
func judgeAndFormatQuert(query map[string][]string, collection *mongo.Collection) (int, int, string, error) {

	//获取query为字符串
	qRange := ":"
	if query["range"] != nil {
		qRangeMap, _ := query["range"]
		if qRangeMap[0] != "" {
			qRange = qRangeMap[0]
		}
	}
	qSort := ""
	if query["sort"] != nil {
		qSortMap, _ := query["sort"]
		qSort = qSortMap[0]
	}

	//判定并格式化
	if strings.Count(qRange, ":") != 1 {
		return -2, -2, "", errors.New("错误的参数：range")
	}
	index := strings.Index(qRange, ":")

	var num = countNum(collection)
	var indexMin = 0
	var indexMax = num
	var err error
	if qRange[:index] != "" {
		indexMin, err = strconv.Atoi(qRange[:index])
		if err != nil {
			return -2, -2, "", errors.New("错误的参数：range")
		}
	}
	if qRange[index+1:] != "" {
		indexMax, err = strconv.Atoi(qRange[index+1:])
		if err != nil {
			return -2, -2, "", errors.New("错误的参数：range")
		}
	}

	//合法性验证
	if indexMin < 0 || indexMax > num || indexMin >= indexMax {
		return -2, -2, "", errors.New("错误的参数：range")
	}
	if qSort != "" && qSort != "D" && qSort != "A" {
		return -2, -2, "", errors.New("错误的参数：sort")
	}

	return indexMin, indexMax, qSort, nil
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

//根据_id查询单张图片
func getImageById(id string) (string, error) {

	//连接
	client, err := connectMongo()
	if err != nil {
		return "", err
	}

	// 指定获取要操作的数据集
	collectionLink := conf.Mongodb.Collection
	collection := client.Database(conf.Mongodb.Db).Collection(collectionLink)
	fmt.Println("Connected to " + conf.Mongodb.Db + "!")

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

//根据_id查询单张图片
func getImageStatusById(id string) ([]byte, error) {

	//连接
	client, err := connectMongo()
	if err != nil {
		return nil, err
	}

	// 指定获取要操作的数据集
	collectionLink := conf.Mongodb.Collection
	collection := client.Database(conf.Mongodb.Db).Collection(collectionLink)
	fmt.Println("Connected to " + conf.Mongodb.Db + "!")

	//按照_id查询图片地址
	result, err := findById(id, collection)
	if err != nil {
		return nil, err
	}

	statusJson, err := json.Marshal(result)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}

	//断开连接
	err = disconnectMongo(client)
	if err != nil {
		return nil, err
	}

	return statusJson, nil
}

//随机获得单张图片
func getImageRandom(qType string) (string, error) {

	//连接
	client, err := connectMongo()
	if err != nil {
		return "", err
	}

	// 指定获取要操作的数据集
	collectionLink := conf.Mongodb.Collection
	collection := client.Database(conf.Mongodb.Db).Collection(collectionLink)
	fmt.Println("Connected to " + conf.Mongodb.Db + "!")

	//获取图库图片数量
	var num int
	rand.Seed(time.Now().UnixNano())
	if qType == "" {
		num = countNum(collection, bson.D{{}})
	} else {
		num = countNum(collection, bson.D{{"info.content", qType}})
	}

	if num == 0 {
		return "TypeWrong", errors.New("No image found in type : " + qType)
	}

	skipNum := rand.Intn(num)

	//搜索图片
	result, err := findSkipNum(skipNum, qType, collection)
	if err != nil {
		return "", err
	}

	//获得地址
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

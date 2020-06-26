package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

//将图片从原地址转移到新地址，插入数据库
func imageInsertAndRemove(oldpath string, imageLibrary string, collection *mongo.Collection) bool {

	//获取MD5，时间，后缀
	imageMd5S := getMd5(oldpath)
	imageTime := time.Now()
	ex := getEx(oldpath)

	//MD5去重
	if isExistMd5(imageMd5S, collection) {
		fmt.Println("md5:" + imageMd5S + "已存在")
		return false
	}

	//移动文件
	newPath := mvFile(oldpath, imageMd5S, imageTime, ex)
	if newPath == "" {
		//同名去重
		return false
	}

	//插入数据库
	var setu setuImage
	setu.Md5 = imageMd5S
	setu.Timestamp = int(imageTime.Unix())
	setu.Info = []imageInfo{{"ex", ex}, {"type", imageLibrary}}
	insertOneOptions := options.InsertOne()
	insertOneMonge(setu, collection, insertOneOptions)
	fmt.Println(oldpath, newPath)

	return true
}

//操作库
func mongodboperation() {

	//连接
	client := connectMongo()

	// 指定获取要操作的数据集
	collectionLink := "setu_image"
	collection := client.Database(mongodb_link.db).Collection(collectionLink)
	fmt.Println("Connected to " + mongodb_link.db + "!")

	// 查询数据
	findOneFilter := bson.D{{}}
	findOneOptions := options.FindOne()
	result, err := findOneMonge(collection, findOneFilter, findOneOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %T\n", result)

	//写入数据
	result.Timestamp = int(time.Now().Unix())
	insertOneOptions := options.InsertOne()
	insertOneMonge(result, collection, insertOneOptions)

	/*
		//差人多条数据
		findFilter :=bson.D{{}}
		findOptions := options.Find()
		findOptions.SetSort(bson.D{{"id", -1}})
		findOptions.SetLimit(1)
		results := findMonge(collection, findFilter, findOptions)
		fmt.Println(results)
	*/

	//断开连接
	disconnectMongo(client)
}

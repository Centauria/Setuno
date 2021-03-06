package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//将图片从原地址转移到新地址，插入数据库
func imageInsertAndRemove(oldpath string, legacyLabel string, collection *mongo.Collection) bool {

	//获取MD5，去重
	imageMd5S := getMd5(oldpath)
	if isExistMd5(imageMd5S, collection) {
		fmt.Println("md5:" + imageMd5S + "已存在")
		return false
	}

	//获取MD5，时间，后缀
	imageTime := time.Now()
	ext := getExt(oldpath)

	//移动文件
	newPath, err := mvFile(oldpath, imageMd5S, imageTime, ext)
	if newPath == "" {
		//同名去重
		fmt.Println(err)
		return false
	}

	//插入数据库
	setu := make(bson.M)
	setu["md5"] = imageMd5S
	setu["timestamp"] = int(imageTime.Unix())
	setu["ext"] = ext
	info := make(bson.M)
	info["legacy_label"] = legacyLabel
	info["dim"] = "2"
	setu["info"] = info
	insertOneOptions := options.InsertOne()
	err = insertOneMonge(setu, collection, insertOneOptions)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(oldpath, newPath)

	return true
}

//操作库
func mongodboperation() error {

	//连接
	client, err := connectMongo()
	if err != nil {
		return err
	}

	// 指定获取要操作的数据集
	collectionLink := conf.Mongodb.Collection
	collection := client.Database(conf.Mongodb.Db).Collection(collectionLink)
	fmt.Println("Connected to " + conf.Mongodb.Db + "!")

	// 查询数据
	findOneFilter := bson.D{{}}
	findOneOptions := options.FindOne()
	result, err := findOneMonge(collection, findOneFilter, findOneOptions)
	if err != nil {
		return err
	}
	fmt.Println(result)

	//写入数据
	/*
		result.Timestamp = int(time.Now().Unix())
		insertOneOptions := options.InsertOne()
		err = insertOneMonge(result, collection, insertOneOptions)
		if err != nil {
			return err
		}
	*/

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
	err = disconnectMongo(client)
	if err != nil {
		return err
	}

	return nil
}

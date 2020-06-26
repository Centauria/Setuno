package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//连接库
func connectMongo() (*mongo.Client, error) {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://" + mongodb_link.user + ":" + mongodb_link.pass + "@" + mongodb_link.host + ":" + mongodb_link.port)

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB!")

	return client, nil
}

//断开连接
func disconnectMongo(client *mongo.Client) error {
	err := client.Disconnect(context.TODO())
	if err != nil {
		return err
	}
	fmt.Println("Connection to MongoDB closed.")
	return nil
}

//读单条数据
func findOneMonge(collection *mongo.Collection, filter interface{}, findOneOptions *options.FindOneOptions) (*setuImage, error) {
	var result setuImage
	err := collection.FindOne(context.TODO(), filter, findOneOptions).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//多条数据查找
func findMonge(collection *mongo.Collection, filter interface{}, findOptions *options.FindOptions) ([]setuImage, error) {
	var results []setuImage
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("Found  document: %+v\n", results)
	return results, nil
}

//写数据
func insertOneMonge(s1 setuImage, collection *mongo.Collection, insertOneOptions *options.InsertOneOptions) error {

	s1.Id = primitive.NewObjectID()
	_, err := collection.InsertOne(context.TODO(), s1, insertOneOptions)
	if err != nil {
		return err
	}
	return nil
}

//统计数量
func countId(collection *mongo.Collection) int {
	findOptions := options.Find()
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return 0
	}
	defer cur.Close(ctx)
	var i = 0
	for cur.Next(ctx) {
		i = i + 1
	}

	return i
}

//检测给定MD5是否存在于数据库
func isExistMd5(md5 string, collection *mongo.Collection) bool {
	findOneFilter := bson.D{{"md5", md5}}
	findOneOptions := options.FindOne()
	_, err := findOneMonge(collection, findOneFilter, findOneOptions)
	if err != nil {
		return false
	}

	return true
}

//根据_id查询数据
func findById(id string, collection *mongo.Collection) (*setuImage, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	findOneFilter := bson.D{{"_id", objectId}}
	findOneOptions := options.FindOne()
	result, err := findOneMonge(collection, findOneFilter, findOneOptions)
	if err != nil {
		return nil, err
	}
	return result, nil
}

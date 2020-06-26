package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type setuImage struct {
	Id        int         `bson:"id"`
	Md5       string      `bson:"md5"`
	Timestamp int         `bson:"timestamp"`
	Info      []imageInfo `bson:"info"`
}

type imageInfo struct {
	Name    string `bson:"name"`
	Content string `bson:"content"`
}

//连接库
func connectMongo() *mongo.Client {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://" + mongodb_link.user + ":" + mongodb_link.pass + "@" + mongodb_link.host + ":" + mongodb_link.port)

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return client
}

//断开连接
func disconnectMongo(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

//读数据
func findoneMonge(collection *mongo.Collection) setuImage {
	var result setuImage
	filter := bson.D{{}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
	return result
}

//写数据
func insertoneMonge(s1 setuImage, collection *mongo.Collection) {

	insertResult, err := collection.InsertOne(context.TODO(), s1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

//操作库
func mongodboperation() {

	//连接
	client := connectMongo()

	// 指定获取要操作的数据集
	collection_link := "setu_image"
	collection := client.Database(mongodb_link.db).Collection(collection_link)
	fmt.Println("Connected to " + mongodb_link.db + "!")

	// 查询数据
	result := findoneMonge(collection)
	fmt.Printf("Found a single document: %T\n", result)
	fmt.Println(result)
	/*
		//写入数据
		result.Id = 2
		insertoneMonge(result, collection)

	*/
	/*
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		cur, err := collection.Find(ctx, bson.D{})
		if err != nil { log.Fatal(err) }
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var result bson.M
			err := cur.Decode(&result)
			if err != nil { log.Fatal(err) }
			// do something with result....
			fmt.Println(result)
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}*/

	//断开连接
	disconnectMongo(client)
}

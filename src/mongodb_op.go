package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

//连接库
func connectMongo() (*mongo.Client, error) {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://" + conf.Mongodb.User + ":" + conf.Mongodb.Pass + "@" + conf.Mongodb.Host + ":" + strconv.Itoa(conf.Mongodb.Port))

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
func findOneMonge(collection *mongo.Collection, filter interface{}, findOneOptions *options.FindOneOptions) (bson.M, error) {
	var result bson.M
	err := collection.FindOne(context.TODO(), filter, findOneOptions).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//多条数据查找
func findMonge(collection *mongo.Collection, filter interface{}, findOptions *options.FindOptions) ([]bson.M, error) {
	var results []bson.M
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
func insertOneMonge(s1 bson.M, collection *mongo.Collection, insertOneOptions *options.InsertOneOptions) error {

	s1["_id"] = primitive.NewObjectID()
	_, err := collection.InsertOne(context.TODO(), s1, insertOneOptions)
	if err != nil {
		return err
	}
	return nil
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
func findById(id string, collection *mongo.Collection) (bson.M, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var result bson.M
	findOneFilter := bson.D{{"_id", objectId}}
	findOneOptions := options.FindOne()
	result, err = findOneMonge(collection, findOneFilter, findOneOptions)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//统计数量
func countNum(collection *mongo.Collection, filter interface{}) int {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	num, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0
	}
	return int(num)
}

//顺序返回某个范围内的_id
func findIdByRangeA(indexMin int, indexMax int, collection *mongo.Collection) ([]string, error) {

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	findOptions := options.Find()
	findOptions.SetSkip(int64(indexMin))
	findOptions.SetLimit(int64(indexMax - indexMin))
	findOptions.SetProjection(bson.D{{"_id", 1}})
	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}

	var ids []string

	result := make(bson.M)
	result["_id"], _ = primitive.ObjectIDFromHex("")
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		err = cur.Decode(&result)
		if err != nil {
			return ids, err
		}
		ids = append(ids, result["_id"].(primitive.ObjectID).Hex())
	}
	return ids, nil
}

//逆序返回某个范围内的_id
func findIdByRangeD(indexMin int, indexMax int, collection *mongo.Collection) ([]string, error) {

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	num := countNum(collection, bson.D{{}})
	findOptions := options.Find()
	findOptions.SetSkip(int64(num - indexMax))
	findOptions.SetLimit(int64(indexMax - indexMin))
	findOptions.SetProjection(bson.D{{"_id", 1}})
	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}

	var ids []string

	result := make(bson.M)
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		err = cur.Decode(&result)
		if err != nil {
			return ids, err
		}
		ids = append(ids, result["_id"].(primitive.ObjectID).Hex())
	}
	for i, j := 0, indexMax-indexMin-1; i < j; i, j = i+1, j-1 {
		ids[i], ids[j] = ids[j], ids[i]
	}

	return ids, nil

}

//获得第skipNum条数据
func findSkipNum(skipNum int, qType string, collection *mongo.Collection) (bson.M, error) {
	findOneFilter := bson.D{{"info.legacy_label", qType}}
	if qType == "" {
		findOneFilter = bson.D{}
	}
	findOneOptions := options.FindOne()
	findOneOptions.SetSkip(int64(skipNum))
	result, err := findOneMonge(collection, findOneFilter, findOneOptions)
	if err != nil {
		return nil, err
	}
	return result, nil
}

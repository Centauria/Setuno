package main

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"strings"
)

//根据get参数获取_id
func getIdByGet(query map[string][]string) ([]string, error) {

	//连接
	client, err := connectMongo()
	if err != nil {
		return nil, err
	}

	// 指定获取要操作的数据集
	collectionLink := conf.Mongodb.Collection
	collection := client.Database(conf.Mongodb.Db).Collection(collectionLink)
	fmt.Println("Connected to " + conf.Mongodb.Db + "!")

	//得到正确query
	indexMin, indexMax, qSort, err := judgeAndFormatQuert(query, collection)
	if err != nil {
		if indexMin == 0 && indexMax == 0 && qSort == "" {
			return []string{}, err
		}
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

	var num = countNum(collection, bson.D{{}})
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
		return 0, 0, "", errors.New("错误的参数：range")
	}
	if qSort != "" && qSort != "D" && qSort != "A" {
		return -2, -2, "", errors.New("错误的参数：sort")
	}

	return indexMin, indexMax, qSort, nil
}

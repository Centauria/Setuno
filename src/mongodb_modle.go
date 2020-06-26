package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type setuImage struct {
	Id        primitive.ObjectID `bson:"_id"`
	Md5       string             `bson:"md5"`
	Timestamp int                `bson:"timestamp"`
	Info      []imageInfo        `bson:"info"`
}

type imageInfo struct {
	Name    string `bson:"name"`
	Content string `bson:"content"`
}

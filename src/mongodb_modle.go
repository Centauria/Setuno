package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type setuImage struct {
	Id        primitive.ObjectID `bson:"_id"`
	Md5       string             `bson:"md5"`
	Timestamp int                `bson:"timestamp"`
	Ext       string             `bson:"ext"`
	Info      imageInfo          `bson:"info"`
}

type imageInfo struct {
	LegacyLabel string `bson:"legacy_label"`
}

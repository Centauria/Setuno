package main

type setuImage struct {
	Md5       string      `bson:"md5"`
	Timestamp int         `bson:"timestamp"`
	Info      []imageInfo `bson:"info"`
}

type imageInfo struct {
	Name    string `bson:"name"`
	Content string `bson:"content"`
}

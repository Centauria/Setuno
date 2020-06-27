package main

import "runtime"

// 配置文件
// file: config.go
type configList struct {
	GoVersion string `json:"GoVersion"`
	Os        string `json:"os"`
	Arch      string `json:"arch"`
	Version   string `json:"version"`
}

var config configList

type mongodb struct {
	host string
	port string
	user string
	pass string
	db   string
}

var mongodbLink mongodb

// 初始化配置
func initConf() {
	config = configList{runtime.Version()[2:], runtime.GOOS, runtime.GOARCH, "0.1"}
	mongodbLink = mongodb{"jinfans.top", "27017", "bot", "bot", "bot"}
}

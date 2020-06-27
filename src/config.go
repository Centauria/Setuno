package main

import (
	"encoding/json"
	"os"
	"runtime"
)

// 配置文件
type configuration struct {
	Program program `json:"program"`
	Info    info    `json:"info"`
	Mongodb mongodb `json:"mongodb"`
	Path    path    `json:"path"`
}

// 程序设置
type program struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// 版本配置
type info struct {
	GoVersion string `json:"go_version"`
	Os        string `json:"os"`
	Arch      string `json:"arch"`
	Version   string `json:"version"`
}

//数据库配置
type mongodb struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	User       string `json:"user"`
	Pass       string `json:"pass"`
	Db         string `json:"db"`
	Collection string `json:"collection"`
}

//路径配置
type path struct {
	ImagePath string `json:"image_path"`
	TempPath  string `json:"temp_path"`
}

var conf configuration

// 初始化配置
func initConf(configPath string) error {

	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&conf)
	if err != nil {
		return err
	}

	conf.Info.GoVersion = runtime.Version()[2:]
	conf.Info.Os = runtime.GOOS
	conf.Info.Arch = runtime.GOARCH

	return nil
}

package main

import "runtime"

// 配置文件
// file: config.go
type config_list struct {
	Go_version string `json:"go_version"`
	Os         string `json:"os"`
	Arch       string `json:"arch"`
	Version    string `json:"version"`
}

var config = config_list{runtime.Version()[2:], runtime.GOOS, runtime.GOARCH, "0.1"}

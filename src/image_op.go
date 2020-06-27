package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// 获取md5
func getMd5(path string) string {
	h := md5.New()
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	io.Copy(h, f)
	return hex.EncodeToString(h.Sum(nil))
}

// 获取后缀
func getEx(path string) string {
	index := strings.LastIndex(path, ".")
	return path[index+1:]
}

//移动文件到新目录，返回新目录
func mvFile(filePath string, imageMd5S string, imageTime time.Time, ex string) (string, error) {

	//TODO:加上Hour
	//获得MD5、时间、后缀
	imageMd5 := []rune(imageMd5S)
	year := strconv.FormatInt(int64(imageTime.Year()), 10)
	month := strconv.FormatInt(int64(imageTime.Month()), 10)
	day := strconv.FormatInt(int64(imageTime.Day()), 10)
	minute := strconv.FormatInt(int64(imageTime.Minute()), 10)
	second := strconv.FormatInt(int64(imageTime.Second()), 10)

	// 根据文件名判断目录是否存在，若不存在，创建目录
	dirPath := conf.Path.ImagePath + year + "/" + month + "/" + string(imageMd5[0]) + "/" + string(imageMd5[1]) + "/" +
		string(imageMd5[2]) + "/" + string(imageMd5[3]) + "/"

	_, err := os.Stat(dirPath)
	if !(err == nil || os.IsExist(err)) {
		//目录不存在，创建目录
		_ = os.MkdirAll(dirPath, 0777)
	}

	//判断文件名是否存在，若存在，返回空
	fileName := imageMd5S[4:8] + day + minute + second + "." + ex
	newPath := dirPath + fileName

	_, err = os.Stat(newPath)
	if err == nil || os.IsExist(err) {
		//目录存在，返回空
		return "", err
	}

	// 将图片剪贴到目录下
	err = os.Rename(filePath, newPath)
	if err != nil {
		return "nil", err
	}

	return newPath, nil
}

//根据记录组织url
func getUrlByResult(result *setuImage) string {

	//TODO:加上Hour
	//时间戳转为时间
	imageTime := time.Unix(int64(result.Timestamp), 0)
	year := strconv.FormatInt(int64(imageTime.Year()), 10)
	month := strconv.FormatInt(int64(imageTime.Month()), 10)
	day := strconv.FormatInt(int64(imageTime.Day()), 10)
	minute := strconv.FormatInt(int64(imageTime.Minute()), 10)
	second := strconv.FormatInt(int64(imageTime.Second()), 10)

	//MD5
	imageMd5S := result.Md5
	imageMd5 := []rune(imageMd5S)

	//ex
	ex := result.Info[0].Content

	dirPath := conf.Path.ImagePath + year + "/" + month + "/" + string(imageMd5[0]) + "/" + string(imageMd5[1]) + "/" +
		string(imageMd5[2]) + "/" + string(imageMd5[3]) + "/"
	fileName := imageMd5S[4:8] + day + minute + second + "." + ex
	newPath := dirPath + fileName

	return newPath

}

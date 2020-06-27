package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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
func getExt(path string) string {
	index := strings.LastIndex(path, ".")
	return path[index+1:]
}

//移动文件到新目录，返回新目录
func mvFile(filePath string, imageMd5S string, imageTime time.Time, ext string) (string, error) {

	//TODO:加上Hour
	//获得MD5、时间、后缀
	imageMd5 := []rune(imageMd5S)
	year := strconv.FormatInt(int64(imageTime.Year()), 10)
	month := changeIntoTwoDigit(strconv.FormatInt(int64(imageTime.Month()), 10))
	day := changeIntoTwoDigit(strconv.FormatInt(int64(imageTime.Day()), 10))
	hour := changeIntoTwoDigit(strconv.FormatInt(int64(imageTime.Hour()), 10))
	minute := changeIntoTwoDigit(strconv.FormatInt(int64(imageTime.Minute()), 10))
	second := changeIntoTwoDigit(strconv.FormatInt(int64(imageTime.Second()), 10))

	// 根据文件名判断目录是否存在，若不存在，创建目录
	dirPath := conf.Path.ImagePath + year + "/" + month + "/" + string(imageMd5[0]) + "/" + string(imageMd5[1]) + "/" +
		string(imageMd5[2]) + "/" + string(imageMd5[3]) + "/"

	_, err := os.Stat(dirPath)
	if !(err == nil || os.IsExist(err)) {
		//目录不存在，创建目录
		_ = os.MkdirAll(dirPath, 0777)
	}

	//判断文件名是否存在，若存在，返回空
	fileName := imageMd5S[4:8] + day + hour + minute + second + "." + ext
	newPath := dirPath + fileName

	_, err = os.Stat(newPath)
	if err == nil || os.IsExist(err) {
		//目录存在，返回空
		return "", err
	}

	// 将图片剪贴到目录下
	err = os.Rename(filePath, newPath)
	if err != nil {
		//return "nil", err
	}

	return newPath, nil
}

//根据记录组织url
func getUrlByResult(result bson.M) string {

	//TODO:加上Hour
	//时间戳转为时间
	imageTime := time.Unix(int64(result["timestamp"].(int32)), 0)
	year := strconv.FormatInt(int64(imageTime.Year()), 10)
	month := changeIntoTwoDigit(strconv.FormatInt(int64(imageTime.Month()), 10))
	day := changeIntoTwoDigit(strconv.FormatInt(int64(imageTime.Day()), 10))
	hour := changeIntoTwoDigit(strconv.FormatInt(int64(imageTime.Hour()), 10))
	minute := changeIntoTwoDigit(strconv.FormatInt(int64(imageTime.Minute()), 10))
	second := changeIntoTwoDigit(strconv.FormatInt(int64(imageTime.Second()), 10))

	//MD5
	imageMd5S := result["md5"].(string)
	imageMd5 := []rune(imageMd5S)

	//ext
	ext := result["ext"].(string)

	dirPath := conf.Path.ImagePath + year + "/" + month + "/" + string(imageMd5[0]) + "/" + string(imageMd5[1]) + "/" +
		string(imageMd5[2]) + "/" + string(imageMd5[3]) + "/"
	fileName := imageMd5S[4:8] + day + hour + minute + second + "." + ext
	newPath := dirPath + fileName

	return newPath

}

func changeIntoTwoDigit(old string) string {
	if len(old) == 1 {
		return "0" + old
	}
	return old
}

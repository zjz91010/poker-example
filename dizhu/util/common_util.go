package util

import (
	"math/rand"
	"strings"
)

//公共方法

/**
随机生成字母+数字
n:生成随机数个数
**/
func RandomString(n int) string {
	letterAndNum := "abcdefghijkmnpqrstuvwxyzABCDEFGHIJKMNPQRSTUVWXYZ23456789"
	//如果不想重复可以使用这个把先引进time包rand,Intn=r.Intn
	// r := rand.New(rand.NewSource(time.Now().Unix()))
	str := make([]byte, n)
	for i := range str {
		str[i] = letterAndNum[rand.Intn(len(letterAndNum))]
	}
	return string(str)
}

/**
拼接方法
每次RandomString()生成字符串再2次拼接
n:生成随机数个数的平方
**/
func Concat(n int) string {
	var str strings.Builder
	for i := 0; i < n; i++ {
		str.WriteString(RandomString(n))
	}
	return str.String()
}

/**
拼接字符串
可变参数
**/
func ConcatString(args ...string) string {
	var str strings.Builder
	for i := range args {
		str.WriteString(args[i])
	}
	return str.String()
}

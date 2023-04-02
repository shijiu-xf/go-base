package utils

import (
	"math/rand"
	"path"
	"runtime"
	"strings"
)

// GetCurrentAbPathByCaller 获取项目的绝对路径
func GetCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

var CHARS = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

// GetRandomStr  生成随机数字和字母字符串 lenNum 长度
func GetRandomStr(lenNum int) string {
	str := strings.Builder{}
	length := len(CHARS)
	for i := 0; i < lenNum; i++ {
		l := CHARS[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}

// GetNumRandomStr  生成随机数字字符串([0~9])lenNum 长度
func GetNumRandomStr(lenNum int) string {
	str := strings.Builder{}
	length := 10
	for i := 0; i < lenNum; i++ {
		str.WriteString(CHARS[52+rand.Intn(length)])
	}
	return str.String()
}

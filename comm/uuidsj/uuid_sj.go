package uuidsj

import (
	"crypto/sha512"
	"encoding/hex"
	"github.com/google/uuid"
)

// NewUUIDv4Str 生成v4版本的UUID的字符串
func NewUUIDv4Str() (uid string) {
	uuidByte, _ := uuid.NewRandom()
	uid = uuidByte.String()
	return uid
}

// MD5 使用sha512加密字符串
func MD5(str string) string {
	newMD5 := sha512.New()
	newMD5.Write([]byte(str))
	return hex.EncodeToString(newMD5.Sum(nil))
}

// NewUUIDv4MD5Str 获取一个sha512加密的UUID
func NewUUIDv4MD5Str() string {
	return MD5(NewUUIDv4Str())
}

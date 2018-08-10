package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"

	"enjoypiano.cn/lynxpro/common/logger"
)

// MD5 generate md5 string
func MD5(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

var (
	alphanum = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	num      = "0123456789"

	// TestRandomValue set random value in unit test
	TestRandomValue = ""
)

// Random get random string by specifying the length
func Random(i int) string {
	if TestRandomValue != "" {
		return TestRandomValue
	}
	bytes := make([]byte, i)
	for {
		rand.Read(bytes)
		for i, b := range bytes {
			bytes[i] = alphanum[b%byte(len(alphanum))]
		}
		return string(bytes)
	}
}

// RandomNum get random number string by specifying the length
func RandomNum(i int) string {
	if TestRandomValue != "" {
		return TestRandomValue
	}
	bytes := make([]byte, i)
	for {
		rand.Read(bytes)
		for i, b := range bytes {
			bytes[i] = num[b%byte(len(num))]
		}
		return string(bytes)
	}
}

// Obj2Str convert Object to string
func Obj2Str(data interface{}) string {
	b, err := json.Marshal(data)
	//logger.Debug("JSON:%s", string(b))
	if err != nil {
		logger.Debug("Failed to convert object to json:%s", err)
		return ""
	}
	return string(b)
}

// Str2Obj convert string to object
func Str2Obj(data string, obj interface{}) {
	if len(data) == 0 {
		return
	}
	if err := json.Unmarshal([]byte(data), obj); err != nil {
		logger.Debug("Failed to convert json to object:%s(%s)", err, data)
	}
}

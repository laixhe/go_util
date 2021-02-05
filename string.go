package goutil

import (
	"math/rand"
	"time"
)

var RandomString = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var RandomStringToLower = []byte("0123456789abcdefghijklmnopqrstuvwxyz")
var RandomStringInt = []byte("0123456789")

// Substr 字符串截取
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

// GetRandomString 生成随机字符串
// l int 随机的长度
// is int 0=不包含大写 1=包含大写 2=只有数字 (默认为 0)
func GetRandomString(l int, is ...int) string {

	bytes := RandomStringToLower
	if len(is) > 0 {
		if is[0] == 1 {
			bytes = RandomString
		}
		if is[0] == 2 {
			bytes = RandomStringInt
		}
	}

	bytesLen := len(bytes)
	result := make([]byte, 0, l)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(bytesLen)])
	}
	return string(result)
}

package goutil

import (
	"math/big"
	"strconv"
	"time"

	cd "crypto/rand"
	md "math/rand"
)

// GetMRandInt64 生成区间随机数
func GetMRandInt64(min, max int64) int64 {
	//time.Sleep(time.Nanosecond)
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return md.New(md.NewSource(time.Now().UnixNano())).Int63n(max-min) + min
}

// GetCRandInt64 生成区间随机数
func GetCRandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	result, _ := cd.Int(cd.Reader, big.NewInt(max-min))
	return result.Int64() + min
}

// StringToInt 字符串转整型
func StringToInt(s string) int {
	if s == "" {
		return 0
	}
	id, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0
	}
	return int(id)
}

// StringToUint 字符串转整型
func StringToUint(s string) uint {
	if s == "" {
		return 0
	}
	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return uint(id)
}

// StringToInt32 字符串转整型
func StringToInt32(s string) int32 {
	if s == "" {
		return 0
	}
	id, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0
	}
	return int32(id)
}

// StringToUint32 字符串转整型
func StringToUint32(s string) uint32 {
	if s == "" {
		return 0
	}
	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return uint32(id)
}

// StringToInt64 字符串转整型
func StringToInt64(s string) int64 {
	if s == "" {
		return 0
	}
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return id
}

// StringToUint64 字符串转整型
func StringToUint64(s string) uint64 {
	if s == "" {
		return 0
	}
	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return id
}

// StringToFloat 字符串转浮点型
func StringToFloat(s string) float64 {
	if s == "" {
		return 0
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

// StringToBool 字符串转浮布尔
func StringToBool(s string) bool {
	if s == "" {
		return false
	}
	b, _ := strconv.ParseBool(s)
	return b
}

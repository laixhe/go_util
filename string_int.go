package goutil

import (
	"errors"
	"strconv"
	"strings"
)

// Int8ArrayToString 拼接 int8 为 string
func Int8ArrayToString(data []int8, sep string) string {

	dataLen := len(data)
	if dataLen == 0 {
		return ""
	}
	dataLen--

	var sepLen = len(sep)
	var str = strings.Builder{}

	for k, v := range data {
		if k != dataLen {
			str.WriteString(strconv.FormatInt(int64(v), 10))
			if sepLen > 0 {
				str.WriteString(sep)
			}
		} else {
			str.WriteString(strconv.FormatInt(int64(v), 10))
		}
	}

	return str.String()
}

// Uint8ArrayToString 拼接 uint8 为 string
func Uint8ArrayToString(data []uint8, sep string) string {

	dataLen := len(data)
	if dataLen == 0 {
		return ""
	}
	dataLen--

	var sepLen = len(sep)
	var str = strings.Builder{}

	for k, v := range data {
		if k != dataLen {
			str.WriteString(strconv.FormatInt(int64(v), 10))
			if sepLen > 0 {
				str.WriteString(sep)
			}
		} else {
			str.WriteString(strconv.FormatInt(int64(v), 10))
		}
	}

	return str.String()
}

// IntArrayToString 拼接 int 为 string
func IntArrayToString(data []int, sep string) string {

	dataLen := len(data)
	if dataLen == 0 {
		return ""
	}
	dataLen--

	var sepLen = len(sep)
	var str = strings.Builder{}

	for k, v := range data {
		if k != dataLen {
			str.WriteString(strconv.Itoa(v))
			if sepLen > 0 {
				str.WriteString(sep)
			}
		} else {
			str.WriteString(strconv.Itoa(v))
		}
	}

	return str.String()
}

// UintArrayToString 拼接 uint 为 string
func UintArrayToString(data []uint, sep string) string {

	dataLen := len(data)
	if dataLen == 0 {
		return ""
	}
	dataLen--

	var sepLen = len(sep)
	var str = strings.Builder{}

	for k, v := range data {
		if k != dataLen {
			str.WriteString(strconv.FormatInt(int64(v), 10))
			if sepLen > 0 {
				str.WriteString(sep)
			}
		} else {
			str.WriteString(strconv.FormatInt(int64(v), 10))
		}
	}

	return str.String()
}

// Int32ArrayToString 拼接 int32 为 string
func Int32ArrayToString(data []int32, sep string) string {

	dataLen := len(data)
	if dataLen == 0 {
		return ""
	}
	dataLen--

	var sepLen = len(sep)
	var str = strings.Builder{}

	for k, v := range data {
		if k != dataLen {
			str.WriteString(strconv.FormatInt(int64(v), 10))
			if sepLen > 0 {
				str.WriteString(sep)
			}
		} else {
			str.WriteString(strconv.FormatInt(int64(v), 10))
		}
	}

	return str.String()
}

// Uint32ArrayToString 拼接 uint32 为 string
func Uint32ArrayToString(data []uint32, sep string) string {

	dataLen := len(data)
	if dataLen == 0 {
		return ""
	}
	dataLen--

	var sepLen = len(sep)
	var str = strings.Builder{}

	for k, v := range data {
		if k != dataLen {
			str.WriteString(strconv.FormatInt(int64(v), 10))
			if sepLen > 0 {
				str.WriteString(sep)
			}
		} else {
			str.WriteString(strconv.FormatInt(int64(v), 10))
		}
	}

	return str.String()
}

// Int64ArrayToString 拼接 int64 为 string
func Int64ArrayToString(data []int64, sep string) string {

	dataLen := len(data)
	if dataLen == 0 {
		return ""
	}
	dataLen--

	var sepLen = len(sep)
	var str = strings.Builder{}

	for k, v := range data {
		if k != dataLen {
			str.WriteString(strconv.FormatInt(v, 10))
			if sepLen > 0 {
				str.WriteString(sep)
			}
		} else {
			str.WriteString(strconv.FormatInt(v, 10))
		}
	}

	return str.String()
}

// Uint64ArrayToString 拼接 uint64 为 string
func Uint64ArrayToString(data []uint64, sep string) string {

	dataLen := len(data)
	if dataLen == 0 {
		return ""
	}
	dataLen--

	var sepLen = len(sep)
	var str = strings.Builder{}

	for k, v := range data {
		if k != dataLen {
			str.WriteString(strconv.FormatUint(v, 10))
			if sepLen > 0 {
				str.WriteString(sep)
			}
		} else {
			str.WriteString(strconv.FormatUint(v, 10))
		}
	}

	return str.String()
}

// StringToInt8Array 分割字符串为 int8 数组
func StringToInt8Array(s, sep string) ([]int8, error) {
	if s == "" {
		return nil, nil
	}

	if sep == "" {
		return nil, errors.New("StringToInt8Array err: sep empty")
	}

	split := strings.Split(s, sep)
	data := make([]int8, 0, len(split))

	for _, v := range split {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("StringToInt8Array err: " + err.Error())
		}

		data = append(data, int8(i))
	}

	return data, nil
}

// StringToUint8Array 分割字符串为 uint8 数组
func StringToUint8Array(s, sep string) ([]uint8, error) {
	if s == "" {
		return nil, nil
	}

	if sep == "" {
		return nil, errors.New("StringToUint8Array err: sep empty")
	}

	split := strings.Split(s, sep)
	data := make([]uint8, 0, len(split))

	for _, v := range split {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("StringToUint8Array err: " + err.Error())
		}

		data = append(data, uint8(i))
	}

	return data, nil
}

// StringToIntArray 分割字符串为 int 数组
func StringToIntArray(s, sep string) ([]int, error) {
	if s == "" {
		return nil, nil
	}

	if sep == "" {
		return nil, errors.New("StringToIntArray err: sep empty")
	}

	split := strings.Split(s, sep)
	data := make([]int, 0, len(split))

	for _, v := range split {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("StringToIntArray err: " + err.Error())
		}

		data = append(data, i)
	}

	return data, nil
}

// StringToUintArray 分割字符串为 uint 数组
func StringToUintArray(s, sep string) ([]uint, error) {
	if s == "" {
		return nil, nil
	}

	if sep == "" {
		return nil, errors.New("StringToUintArray err: sep empty")
	}

	split := strings.Split(s, sep)
	data := make([]uint, 0, len(split))

	for _, v := range split {
		i, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return nil, errors.New("StringToUintArray err: " + err.Error())
		}

		data = append(data, uint(i))
	}

	return data, nil
}

// StringToInt32Array 分割字符串为 int32 数组
func StringToInt32Array(s, sep string) ([]int32, error) {
	if s == "" {
		return nil, nil
	}

	if sep == "" {
		return nil, errors.New("StringToInt32Array err: sep empty")
	}

	split := strings.Split(s, sep)
	data := make([]int32, 0, len(split))

	for _, v := range split {
		i, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return nil, errors.New("StringToInt32Array err: " + err.Error())
		}

		data = append(data, int32(i))
	}

	return data, nil
}

// StringToUint32Array 分割字符串为 uint32 数组
func StringToUint32Array(s, sep string) ([]uint32, error) {
	if s == "" {
		return nil, nil
	}

	if sep == "" {
		return nil, errors.New("StringToUint32Array err: sep empty")
	}

	split := strings.Split(s, sep)
	data := make([]uint32, 0, len(split))

	for _, v := range split {
		i, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return nil, errors.New("StringToUint32Array err: " + err.Error())
		}

		data = append(data, uint32(i))
	}

	return data, nil
}

// StringToInt64Array 分割字符串为 int64 数组
func StringToInt64Array(s, sep string) ([]int64, error) {
	if s == "" {
		return nil, nil
	}

	if sep == "" {
		return nil, errors.New("StringToInt64Array err: sep empty")
	}

	split := strings.Split(s, sep)
	data := make([]int64, 0, len(split))

	for _, v := range split {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, errors.New("StringToInt64Array err: " + err.Error())
		}

		data = append(data, i)
	}

	return data, nil
}

// StringToUint64Array 分割字符串为 uint64 数组
func StringToUint64Array(s, sep string) ([]uint64, error) {
	if s == "" {
		return nil, nil
	}

	if sep == "" {
		return nil, errors.New("StringToUint64Array err: sep empty")
	}

	split := strings.Split(s, sep)
	data := make([]uint64, 0, len(split))

	for _, v := range split {
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return nil, errors.New("StringToUint64Array err: " + err.Error())
		}

		data = append(data, i)
	}

	return data, nil
}

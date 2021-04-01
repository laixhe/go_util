package goini

import (
	"log"
	"strings"

	"github.com/go-ini/ini"
)

var conf *ini.File

// Open 加载配置文件
func Open(c string) error {
	var err error
	conf, err = ini.Load(c)
	if err != nil {
		log.Printf("init config err: %v\n", err)
	}
	return err
}

// String 获取单个配置
func String(str string) string {
	strArr := strings.Split(str, ".")
	if len(strArr) == 2 {
		return conf.Section(strArr[0]).Key(strArr[1]).String()
	}
	return conf.Section("").Key(strArr[0]).String()
}

// MustString 获取单个配置 String 并有默认值
func MustString(str, value string) string {
	strArr := strings.Split(str, ".")
	if len(strArr) == 2 {
		return conf.Section(strArr[0]).Key(strArr[1]).MustString(value)
	}
	return conf.Section("").Key(strArr[0]).MustString(value)
}

// Int 获取单个配置
func Int(str string) (int, error) {
	strArr := strings.Split(str, ".")
	if len(strArr) == 2 {
		return conf.Section(strArr[0]).Key(strArr[1]).Int()
	}
	return conf.Section("").Key(strArr[0]).Int()
}

// MustInt 获取单个配置 Int 并有默认值
func MustInt(str string, value int) int {
	strArr := strings.Split(str, ".")
	if len(strArr) == 2 {
		return conf.Section(strArr[0]).Key(strArr[1]).MustInt(value)
	}
	return conf.Section("").Key(strArr[0]).MustInt(value)

}

// Float64 获取单个配置
func Float64(str string) (float64, error) {
	strArr := strings.Split(str, ".")
	if len(strArr) == 2 {
		return conf.Section(strArr[0]).Key(strArr[1]).Float64()
	}
	return conf.Section("").Key(strArr[0]).Float64()
}

// MustFloat64 获取单个配置 Float64 并有默认值
func MustFloat64(str string, value float64) float64 {
	strArr := strings.Split(str, ".")
	if len(strArr) == 2 {
		return conf.Section(strArr[0]).Key(strArr[1]).MustFloat64(value)
	}
	return conf.Section("").Key(strArr[0]).MustFloat64(value)
}

// Bool 获取单个配置
func Bool(str string) (bool, error) {
	strArr := strings.Split(str, ".")
	if len(strArr) == 2 {
		return conf.Section(strArr[0]).Key(strArr[1]).Bool()
	}
	return conf.Section("").Key(strArr[0]).Bool()
}

// MustBool 获取单个配置 Bool 并有默认值
func MustBool(str string, value bool) bool {
	strArr := strings.Split(str, ".")
	if len(strArr) == 2 {
		return conf.Section(strArr[0]).Key(strArr[1]).MustBool(value)
	}
	return conf.Section("").Key(strArr[0]).MustBool(value)
}

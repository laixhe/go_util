package goutil

import (
	"log"
	"strings"

	"github.com/go-ini/ini"
)

var conf *ini.File

// OpenConfig 加载配置文件
func OpenConfig(c string) error {

	var err error
	conf, err = ini.Load(c)
	if err != nil {
		log.Printf("init config err: %v\n", err)
	}

	return err
}

// ConfigString 获取单个配置 String
func ConfigString(str string) string {

	strArr := strings.Split(str, ".")

	if len(strArr) == 2 {
		return conf.Section(strArr[0]).Key(strArr[1]).String()
	}

	return conf.Section("").Key(strArr[0]).String()
}

// ConfigMustString 获取单个配置 String，并有默认值
func ConfigMustString(str, value string) string {

	strArr := strings.Split(str, ".")

	if len(strArr) == 2 {
		return conf.Section(strArr[0]).Key(strArr[1]).MustString(value)
	}

	return conf.Section("").Key(strArr[0]).MustString(value)
}

// ConfigInt 获取单个配置 Int
func ConfigInt(str string) (int, error) {

	strArr := strings.Split(str, ".")

	if len(strArr) == 2 {
		return conf.Section(strArr[0]).Key(strArr[1]).Int()
	}

	return conf.Section("").Key(strArr[0]).Int()
}

// ConfigMustInt 获取单个配置 Int，并有默认值
func ConfigMustInt(str string, value int) int {

	strArr := strings.Split(str, ".")

	if len(strArr) == 2 {
		return conf.Section(strArr[0]).Key(strArr[1]).MustInt(value)
	}

	return conf.Section("").Key(strArr[0]).MustInt(value)

}

// ConfigFloat64 获取单个配置 Float64
func ConfigFloat64(str string) (float64, error) {

	strArr := strings.Split(str, ".")

	if len(strArr) == 2 {
		return conf.Section(strArr[0]).Key(strArr[1]).Float64()
	}

	return conf.Section("").Key(strArr[0]).Float64()
}

// ConfigMustFloat64 获取单个配置 Float64，并有默认值
func ConfigMustFloat64(str string, value float64) float64 {

	strArr := strings.Split(str, ".")

	if len(strArr) == 2 {
		return conf.Section(strArr[0]).Key(strArr[1]).MustFloat64(value)
	}

	return conf.Section("").Key(strArr[0]).MustFloat64(value)
}

// ConfigBool 获取单个配置 Bool
func ConfigBool(str string) (bool, error) {

	strArr := strings.Split(str, ".")

	if len(strArr) == 2 {
		return conf.Section(strArr[0]).Key(strArr[1]).Bool()
	}

	return conf.Section("").Key(strArr[0]).Bool()

}

// ConfigMustBool 获取单个配置 Bool，并有默认值
func ConfigMustBool(str string, value bool) bool {

	strArr := strings.Split(str, ".")

	if len(strArr) == 2 {
		return conf.Section(strArr[0]).Key(strArr[1]).MustBool(value)
	}

	return conf.Section("").Key(strArr[0]).MustBool(value)

}

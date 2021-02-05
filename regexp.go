package goutil

import "regexp"

var (
	emailExp  = regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)
	mobileExp = regexp.MustCompile(`^1[0-9]{10}$`)
	numberExp = regexp.MustCompile(`^[0-9]+$`)
)

// MatchEmail 是否为邮箱
func MatchEmail(s string) bool {
	if s == "" {
		return false
	}
	return emailExp.MatchString(s)
}

// MatchMobile 是否为手机号
func MatchMobile(s string) bool {
	if s == "" {
		return false
	}
	return mobileExp.MatchString(s)
}

// MatchNumber 是否为数整数
func MatchNumber(s string) bool {
	if s == "" {
		return false
	}
	return numberExp.MatchString(s)
}

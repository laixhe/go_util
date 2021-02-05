package goutil

import (
	"fmt"
	"testing"
)

func TestSubstr(t *testing.T) {
	str := Substr("123456789", 0, 6)
	fmt.Println(str)
}

func TestGetRandomString(t *testing.T) {
	str := GetRandomString(20, 1)
	fmt.Println(str, len(str))
}

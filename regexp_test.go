package goutil

import (
	"fmt"
	"testing"
)

func TestMatchEmail(t *testing.T) {
	fmt.Println(MatchEmail("laixhe@laixhe.com"))
	fmt.Println(MatchEmail("laixhe#laixhe.cn"))
}

func TestMatchMobile(t *testing.T) {
	fmt.Println(MatchMobile("13000000000"))
	fmt.Println(MatchMobile("138001380000"))
}

func TestMatchNumber(t *testing.T) {
	fmt.Println(MatchNumber("123"))
	fmt.Println(MatchNumber("a123"))
}
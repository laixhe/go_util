package goutil

import (
	"fmt"
	"testing"
)

func TestGetMRandInt64(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(GetMRandInt64(1000, 9999))
	}
}

func TestGetCRandInt64(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(GetCRandInt64(1000, 9999))
	}
}

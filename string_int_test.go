package goutil

import (
	"fmt"
	"testing"
)

func TestIntArrayToString(t *testing.T) {
	fmt.Println(IntArrayToString([]int{1, 2, 3, 5}, ","))
}

func TestInt32ArrayToString(t *testing.T) {
	fmt.Println(Int32ArrayToString([]int32{1, 2, 3, 5, 4}, ""))
}

func TestInt64ArrayToString(t *testing.T) {
	fmt.Println(Int64ArrayToString([]int64{1, 2, 3, 5, 4, 6}, "|"))
}

func TestStringToIntArray(t *testing.T) {
	fmt.Println(StringToIntArray("12,23,45,2147483647,2147483648", ","))
}

func TestStringToInt32Array(t *testing.T) {
	fmt.Println(StringToInt32Array("12,23,45,56,2147483647,2147483648", ","))
}

func TestStringToInt64Array(t *testing.T) {
	fmt.Println(StringToInt64Array("12,23,45,56,78,2147483647,2147483648", ","))
}

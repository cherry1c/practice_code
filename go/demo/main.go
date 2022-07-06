package main

import (
	"fmt"
	"reflect"
	"strings"
)

func uint64ToString(s []uint64) string {
	fmt.Printf("%v\n", s)
	return fmt.Sprintf("%v", s)
}

func testUint64ToString() {
	s := []uint64{1, 2, 3, 4, 5}
	fmt.Printf("%s\n", uint64ToString(s))
}

func testReflect() {
	s1 := []uint64{1, 2, 3, 4}
	s2 := []uint64{1, 2, 5, 4}
	fmt.Printf("%v\n", reflect.DeepEqual(s1, s2))
}

func PrintSlice() {
	s := "hello world"
	fmt.Printf("%s\n", s[1:3])
}

func Addstring() {
	s := "hello world"
	var b []rune
	for _, val := range s {
		b = append(b, val)
	}
	fmt.Printf("%s\n", string(b))
	strings.Join(b)
}

func main() {
	Addstring()
}

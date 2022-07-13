package main

import (
	"fmt"
	"reflect"
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
	// strings.Join(b)
}

func JudgeMap() {
	m := make(map[int][]int)
	m[1] = []int{1, 2, 3}
	val, ok := m[0]
	if ok {
		fmt.Println("ok ", "val ", val)
	} else {
		fmt.Println("not ok ", "val ", len(val))
	}
}

func IsContainItem(v uint64, s []uint64) bool {
	for _, val := range s {
		if v == val {
			return true
		}
	}
	return false
}

func DelSlice() {
	s := []uint64{1, 2, 3, 4, 5}
	hide := []uint64{12, 2, 3, 4}
	fmt.Println(s)
	for i := 0; i < len(s); {
		if IsContainItem(s[i], hide) {
			s = append(s[:i], s[i+1:]...)
		} else {
			i++
		}
	}
	fmt.Println(s)

}

func testSlice() {
	s := make([]int, 10)
	s = append(s, 1)
	s[2] = 10
	fmt.Println(len(s))
}

func main() {
	testSlice()
}

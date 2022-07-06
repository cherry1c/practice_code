package main

import "fmt"

func test01() {
	s := "hello"
	fmt.Printf("type: %T\n", s[0])
}

func main() {
	longestPalindrome("cbbd")
}

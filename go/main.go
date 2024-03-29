package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

func test01() {
	s := "hello"
	fmt.Printf("type: %T\n", s[0])
}

type MyError struct {
}

func (ei MyError) Error() string {
	return fmt.Sprintf("%#v", ei)
}

func testAtomic() error {
	var filterError atomic.Value
	err1 := errors.New("is err1")
	// err2 := MyError{}
	err3 := fmt.Errorf("%d\n", 100)
	// return err2
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		filterError.Store(err1)
		if err := filterError.Load(); err != nil {
			fmt.Printf("%v\n", err.(error).Error())
		}
	}()
	go func() {
		defer wg.Done()
		filterError.Store(err3)
		if err := filterError.Load(); err != nil {
			fmt.Printf("%v\n", err.(error).Error())
		}
	}()
	wg.Wait()
	if err := filterError.Load(); err != nil {
		fmt.Printf("%v\n", err.(error).Error())
	}
	return nil
}

type Request struct {
	Scene string
}

func readJson() {
	data := `{
    "scene": "test"
}
`
	r := Request{}
	if err := json.Unmarshal([]byte(data), &r); err != nil {
		fmt.Println("unmarshal failed.")
		return
	}
	fmt.Println(r)
}

func main() {
	readJson()

}

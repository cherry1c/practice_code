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

type RType struct {
}

func (s *RType) registerService(a int, b string) int {
	fmt.Println(a, b)
	return a
}

func registerSrv(typeOf reflect.Type, valueOf reflect.Value, paramList []reflect.Value) {
	//ret := valueOf.Call(paramList)
	fmt.Println(valueOf.NumMethod())
	//m := typeOf.Method(0)
	// ret := valueOf.Func.Call(paramList)
	// fmt.Println(ret[0])
}
func testRegisterService() {
	typeOf := reflect.TypeOf(&RType{})
	valueOf := reflect.ValueOf(&RType{})
	paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf("hello world")}
	registerSrv(typeOf, valueOf, paramList)
}

type Abstract interface {
	GetName() string
}

type Instance struct {
}

func (i *Instance) GetName() string {
	return "instance"
}

func test01() {
	var t Abstract
	t = &Instance{}
	t.GetName()
}

type HelloService struct {
	Name string
}

func (h *HelloService) Say() string {
	return "hello " + h.Name
}

func (h HelloService) Start() string {
	return "hello world"
}

func ServiceClient(name string) *HelloService {
	return &HelloService{Name: name}
}

func registerServiceDemo() {
	var registerFunc = reflect.ValueOf(ServiceClient)
	var paramList = reflect.ValueOf("gene")

	t := registerFunc.Call([]reflect.Value{paramList})
	for i := 0; i < t[0].NumMethod(); i++ {
		v1 := t[0].Method(i).Call([]reflect.Value{})
		fmt.Println(v1[0].Interface().(string))
	}
}

func main() {
	registerServiceDemo()
}

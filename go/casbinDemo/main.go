package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
)

func main() {
	e, err := casbin.NewEnforcer("./model.conf")
	//e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		// 处理err
		fmt.Println("failed 12 ", err.Error())
		return
	}

	sub := "alice" // 想要访问资源的用户。
	obj := "data1" // 将被访问的资源。
	act := "read"  // 用户对资源执行的操作。

	//act = "write"
	//e.AddPermissionForUser("alice", "data1", "write")
	e.AddRoleForUser("alice", "data_admin")
	e.AddPolicy("data_admin", "data1", "read")
	e.AddNamedMatchingFunc("g", "data1/*", util.KeyMatch2)

	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		// 处理err
		fmt.Println("failed 20 ", err.Error())
		return
	}

	if ok == true {
		// 允许alice读取data1
		fmt.Println("ok")
	} else {
		// 拒绝请求，抛出异常
		fmt.Println("no permission")
		return
	}

	// 您可以使用BatchEnforce()来批量执行一些请求
	// 这个方法返回布尔切片，此切片的索引对应于二维数组的行索引。
	// 例如results[0] 是{"alice", "data1", "read"}的结果
	//results, err := e.BatchEnforce([[] []interface{}{"alice", "data1", "read"}, {"bob", "datata2", "write"}, {"jack", "data3", "read"}})
	//
	//roles, err := e.GetRolesForUser(sub)
	//if err != nil {
	//	// 处理err
	//	fmt.Println("failed 40 ", err.Error())
	//	return
	//}
	//fmt.Println("roles ", roles)
}

package main

import (
	"fmt"
	"reflect"
)

// 通过反射，修改，num int的值，修改student的值
func test01(b interface{}) {
	//获取reflect.Value
	rVal := reflect.ValueOf(b) //此时rVal为指针类型，无法使用其方法,要用Elem()方法
	rVal.Elem().SetInt(20)
}
func main() {
	var num int = 10
	test01(&num)
	fmt.Println("num =", num)
}

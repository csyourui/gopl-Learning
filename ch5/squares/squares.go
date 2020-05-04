package main

import "fmt"
//squares返回一个函数，函数不接受参数，返回一个整形
func squares() func() int{
	var x int //初始化为0
	return func() int{
		x++
		return x * x
	}
}
func main() {
	f:= squares()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
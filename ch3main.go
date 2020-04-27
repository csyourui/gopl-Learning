package main

import (
	"fmt"
	"gopl/ch3"
)

func main() {
	fmt.Println(ch3.Comma("12131231232"))
	fmt.Println(ch3.Basename("/yourui/go/src/Learn/main.go"))
	fmt.Println(ch3.IntsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}

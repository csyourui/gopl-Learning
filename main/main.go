// package main

// import (
// 	"fmt"
// 	"gopl/ch3"
// 	"gopl/ch4"
// )

// //gopl/ch3/comma.go
// func testCh3Comma() {
// 	fmt.Println(ch3.Comma("12131231232"))
// }

// //gopl/ch3/basename.go
// func testCh3Basename() {
// 	fmt.Println(ch3.Basename("/yourui/go/src/Learn/main.go"))
// }

// //gopl/ch3/intsToString
// func testCh3IntsToString() {
// 	fmt.Println(ch3.IntsToString([]int{1, 2, 3})) // "[1, 2, 3]"
// }

// //gopl/ch3/netflag.go
// func testCh3Netflag() {
// 	var v ch3.Flags = ch3.FlagMulticast | ch3.FlagUp
// 	fmt.Printf("%b %t\n", v, ch3.IsUp(v)) // "10001 true"
// 	ch3.TurnDown(&v)
// 	fmt.Printf("%b %t\n", v, ch3.IsUp(v)) // "10000 false"
// 	ch3.SetBroadcast(&v)
// 	fmt.Printf("%b %t\n", v, ch3.IsUp(v))   // "10010 false"
// 	fmt.Printf("%b %t\n", v, ch3.IsCast(v)) // "10010 true"
// }

// //gopl/ch4/sha256
// func testCh4Sha256() {
// 	s1 := "you"
// 	s2 := "haha"
// 	ch4.Testha256(s1)
// 	fmt.Printf("两个SHA256哈希码中不同bit的数目为：%d\n", ch4.CompareSha256(s1, s2))
// }

// //gopl/ch4/rev
// func testCh4Rev() {
// 	a := []int{0, 1, 2, 3, 4, 5}
// 	ch4.Reverse(a)
// 	fmt.Println(a)

// }

// //gopl/ch4/append.go
// func testCh4Append() {
// 	var x, y []int
// 	for i := 0; i < 10; i++ {
// 		y = ch4.AppendInt(x, i)
// 		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
// 		x = y
// 	}

// 	var a []int
// 	a = append(a, 1)
// 	a = append(a, 2, 3)
// 	a = append(a, 4, 5, 6)
// 	a = append(a, a...) // append the slice x
// 	fmt.Println(a)      // "[1 2 3 4 5 6 1 2 3 4 5 6]"
// }

// //gopl/ch4/nonempty
// func testCh4Nonempty() {
// 	data := []string{"one", "", "three"}
// 	fmt.Printf("%q\n", ch4.Nonempty(data)) // `["one" "three"]`
// 	fmt.Printf("%q\n", data)               // `["one" "three" "three"]`
// }

// //gopl/ch4/dedup
// func testCh4Dedup() {
// 	ch4.Dedup()
// }

// //gopl/ch4/charcount
// func testCh4Charcount() {
// 	ch4.Charcount()
// }

// //gopl/ch4/graph
// func testCh4Graph() {
// 	/*
// 		a - b - d
// 		|	    |
// 		c - e	f
// 	*/
// 	ch4.TestGraph()
// }

// //gopl/ch4/employee_struct
// func testCh4EmployeeStruct() {
// 	ch4.TestEmployeeStruct()
// }

// //gopl/ch4/embed
// func testCh4Embed() {
// 	ch4.TestEmbed()
// }

// //gopl/ch4/movie
// func testCh4Movie() {
// 	ch4.TestMovie()
// }
// func main() {
// 	testCh4Movie()
// }

package ch4

import (
	"crypto/sha256"
	"fmt"
)

//Testha256 print sha256(x) and sha256(X)
func Testha256(s string) {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("sha256(%s) = %x\n", s, sha256.Sum256([]byte(s)))
	fmt.Printf("c1 = sha256(x) = %x\n", c1)
	fmt.Printf("c2 = sha256(X) = %x\n", c2)
	fmt.Printf("c1 == c2 = %t\n", c1 == c2)
	fmt.Printf("Type of c1 is: %T\n", c1)
}

//计算a和b不同的bit位数
func compareInteger(a, b uint32) int {
	c := a ^ b
	res := 0
	for c > 0 {
		res++
		c = c & (c - 1)
	}
	return res
}

//CompareSha256 4.1计算两个SHA256哈希码中不同bit的数目
func CompareSha256(str1, str2 string) int {
	res := 0
	c1 := sha256.Sum256([]byte(str1))
	c2 := sha256.Sum256([]byte(str2))
	for i := 0; i < 10; i++ {
		res += compareInteger(uint32(c1[i]), uint32(c2[i]))
	}
	return res
}

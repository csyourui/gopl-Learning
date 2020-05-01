package ch4

import (
	"bufio"
	"fmt"
	"os"
)

//Dedup 读取多行输入，但是只打印第一次出现的行
func Dedup() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "debup: %v\n", err)
		os.Exit(1)
	}
}

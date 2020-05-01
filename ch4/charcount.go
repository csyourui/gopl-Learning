package ch4

import (
	"bufio"
	"os"
	"unicode/utf8"
)

//Charcount 统计输入中每个Unicode码点出现的次数
func Charcount() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
}

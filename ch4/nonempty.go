package ch4

//Nonempty 在原有slice内存空间之上返回不包含空字符串的字符串
func Nonempty(str []string) []string {
	index := 0
	for _, s := range str {
		if s != "" {
			str[index] = s
			index++
		}
	}
	return str[:index]
}

/*
 * @lc app=leetcode.cn id=884 lang=golang
 *
 * [884] 两句话中的不常见单词
 */

// @lc code=start
func uncommonFromSentences(A string, B string) []string {
	var res []string
	Aset, Bset := strings.Split(A, " "),strings.Split(B, " ")
	MapA, MapB := make(map[string]int),make(map[string]int)
	for _, v := range Aset {
		MapA[v]++
	}
	for _, v := range Bset {
		MapB[v]++
	}
	for _, v := range Aset {
		if _, ok := MapB[v]; !ok && MapA[v] == 1{
			res = append(res, v)
		}
	}
	for _, v := range Bset {
		if _, ok := MapA[v]; !ok && MapB[v] == 1{
			res = append(res, v)
		}
	}
	return res
}
// @lc code=end
/*
 * @lc app=leetcode.cn id=1337 lang=golang
 *
 * [1337] 方阵中战斗力最弱的 K 行
 */

// @lc code=start
func kWeakestRows(mat [][]int, k int) []int {
    sliceSum := func (s []int) int {
		var sum int
		for _, v := range s {
			sum += v
		}
		return sum
	}
	var s [][]int
	for i, v :=  range mat {
		tmp := []int{sliceSum(v), i}
		s = append(s, tmp)
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i][0] < s[j][0] || (s[i][0] == s[j][0] && s[i][1] < s[j][1])
	})
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = s[i][1]
	}
    return res
}
// @lc code=end


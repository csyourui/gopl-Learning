/*
 * @lc app=leetcode.cn id=466 lang=golang
 *
 * [466] 统计重复个数
 */

// @lc code=start
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	i, j, res, L1, L2 := 0, 0, 0, len(s1), len(s2)
	for i < n1 * L1{
		if s1[i % L1] == s2[j % L2] {
			j++
			if j % (n2 * L2) == 0{
				res++
			} 
		}
		i++
	}
	return res
}
// @lc code=end


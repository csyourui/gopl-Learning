/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	var res byte
	for _, c := range s+t {
		res ^= byte(c)
	}
	return res
}
// @lc code=end


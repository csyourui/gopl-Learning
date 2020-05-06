/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func tranform(n int) int{
	res := 0
	for n > 0{
		res += (n%10)*(n%10)
		n /= 10
	}
	return res
}
func isHappy(n int) bool {
	archive := make(map[int]bool)
	for archive[n] == false{
		archive[n] = true
		if n == 1{
			return true
		}
		n = tranform(n)
	}
	return false
}
// @lc code=end


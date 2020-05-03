/*
 * @lc app=leetcode.cn id=1317 lang=golang
 *
 * [1317] 将整数转换为两个无零整数的和
 */

// @lc code=start
func hasZero(n int) bool{
	for n > 0 {
		if n % 10 == 0 {return true}
		n /= 10
	}
	return false
}
func getNoZeroIntegers(n int) []int {
	var res []int
	for i := 1; i <= n/2; i++ {
		if !hasZero(i) && !hasZero(n-i){
			res = append(res, i)
			res = append(res, n-i)
			break
		}
	}
	return res
}
// @lc code=end


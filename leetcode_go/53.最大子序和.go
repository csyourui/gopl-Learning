/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
func maxSubArray(nums []int) int {
	count , res := 0, math.MinInt32
	for _, val := range nums {
		if count < 0{
			count = val
		}else{
			count += val
		}
		res = Max(res, count)
	}
	return res
}
// @lc code=end


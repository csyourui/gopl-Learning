/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func reverseArr(i,j int, nums []int){
	for i < j{
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
func rotate(nums []int, k int)  {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}
	reverseArr(0, n-k-1, nums)
	reverseArr(n-k, n-1, nums)
	reverseArr(0, n-1, nums)
}
// @lc code=end


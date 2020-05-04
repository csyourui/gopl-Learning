/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
func jump(nums []int) int {
	res, maxDis, end := 0, 0, 0
	for i, val := range nums {
		if i == len(nums)-1{
			break
		}
		maxDis = Max(maxDis, i + val)
		if i == end{
			res++
			end = maxDis
		}
	}
	return res
}
// @lc code=end


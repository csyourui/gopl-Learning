/*
 * @lc app=leetcode.cn id=1287 lang=golang
 *
 * [1287] 有序数组中出现次数超过25%的元素
 */

// @lc code=start
func findSpecialInteger(arr []int) int {	
	i, j, L := 0, 0, len(arr)
	F := func(i, j, L int) bool{
		if 4 * (j - i) > L{
			return true
		}
		return false
	}

	for ; i < L && j < L; j++ {
		if arr[j] != arr[i]{
			if F(i, j, L) {
				return arr[i]
			}
			i = j
		}
	}
	if F(i, j, L) {
		return arr[i]
	}
	return -1
}
// @lc code=end


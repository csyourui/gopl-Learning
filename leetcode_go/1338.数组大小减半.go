/*
 * @lc app=leetcode.cn id=1338 lang=golang
 *
 * [1338] 数组大小减半
 */

// @lc code=start
func minSetSize(arr []int) int {
	M := make(map[int]int)
	res, sum := 0, 0
	for _, v := range arr {
		M[v]++
	}
	var count []int
	for _, v := range M {
		count = append(count, v)
	}
	sort.Slice(count, func(i, j int) bool {
		return count[i] > count[j]
	})
	for _, v := range count {
		if sum >= len(arr)/2{
			break
		}
		res++
		sum += v
	}
	return res
}
// @lc code=end


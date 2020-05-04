/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 */

// @lc code=start
func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, v := range nums {
		strs[i] = strconv.Itoa(v)
	}
	sort.Slice(strs, func(i, j int)bool{
		return strs[i]+strs[j] > strs[j]+strs[i]
	})
	res := strings.Join(strs, "")
	if res[0] == '0' {
		return "0"
	}
	return res
}
// @lc code=end


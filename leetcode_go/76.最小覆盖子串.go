/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s, t string) string {
	left, right, start, match := 0, 0, 0, 0
	minLen := math.MaxInt64
	need, window := make(map[rune]int), make(map[rune]int)
	for _, c := range t {
		need[c]++
	}
	for right < len(s) {
		c1 := rune(s[right])
		if need[c1] > 0 {
			window[c1]++
			if window[c1] == need[c1] {
				match++
			}
		}
		right++
		for match == len(need) {
			if right-left < minLen {
				minLen = right - left
				start = left
			}
			c2 := rune(s[left])
			if need[c2] > 0 {
				window[c2]--
				if window[c2] < need[c2] {
					match--
				}
			}
			left++
		}
	}
	if minLen == math.MaxInt64 {
		return ""
	}
	return s[start : start+minLen]
}
// @lc code=end


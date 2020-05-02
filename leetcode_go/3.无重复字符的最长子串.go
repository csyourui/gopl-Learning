/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
func lengthOfLongestSubstring(s string) int {
    M := make(map[rune]int)
    if s == ""{
        return 0
    }
    res, start := 0, 0
    for i, c := range s{
        start = Max(start, M[c])
        res = Max(res, i - start + 1)
        M[c] = i + 1
    }
    return res
}
// @lc code=end


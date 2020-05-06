/*
 * @lc app=leetcode.cn id=583 lang=golang
 *
 * [583] 两个字符串的删除操作
 */

// @lc code=start
func Max(a, b int)int{
	if a >= b {
		return a
	}else{
		return b
	}
}
func minDistance(word1 string, word2 string) int {
	m, n, maxLen := len(word1), len(word2), 0
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1]{
				dp[i][j] = dp[i-1][j-1] + 1
			}else{
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
			maxLen = Max(maxLen, dp[i][j])
		}
	}
	return m+n-2*maxLen
}
// @lc code=end


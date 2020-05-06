/*
 * @lc app=leetcode.cn id=983 lang=golang
 *
 * [983] 最低票价
 */

// @lc code=start
func Min(x, y int) int{
	if x > y{
		return y
	}else{
		return x
	}
}
func mincostTickets(days []int, costs []int) int {
	dp := make([]int, 366+30)
	for i := 1; i < 366; i++ {
		dp[i] = math.MaxInt32;
	}
	j := len(days) - 1
	for i := 365; i >= 1; i-- {
		if j >= 0 && i == days[j]{
			dp[i] = Min(dp[i+1] + costs[0], dp[i+7] + costs[1])
			dp[i] = Min(dp[i], dp[i+30] + costs[2])
			j--
		}else{
			dp[i] = dp[i+1]
		}
	}
	return dp[1]
}
// @lc code=end


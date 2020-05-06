/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func Helper(root *TreeNode, lower, upper int) bool{
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return Helper(root.Left, lower, root.Val) && Helper(root.Right, root.Val, upper)
}
func isValidBST(root *TreeNode) bool {
	return Helper(root, math.MinInt64, math.MaxInt64)
}
// @lc code=end


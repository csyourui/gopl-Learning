/*
 * @lc app=leetcode.cn id=1022 lang=golang
 *
 * [1022] 从根到叶的二进制数之和
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
func dfs(root* TreeNode, curr int, res *int){
	if root.Left != nil{
		dfs(root.Left, root.Val + 2 * curr, res)
	}
	if root.Right != nil{
		dfs(root.Right, root.Val + 2 * curr, res)
	}
	if root.Right == nil && root.Left == nil{
		fmt.Println(root.Val + 2 * curr)
		*res += (root.Val + 2 * curr) % (1000000007)
		*res %= 1000000007
		return
	}
}
func sumRootToLeaf(root *TreeNode) int {
	var res int
	dfs(root, 0, &res)
	return res
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
type NumArray struct {
	Sum []int
}


func Constructor(nums []int) NumArray {
	var res NumArray
	for _, v := range nums {
		if len(res.Sum)	== 0{
			res.Sum = append(res.Sum, v)
		}else{
			res.Sum = append(res.Sum, res.Sum[len(res.Sum)-1] + v)
		}
	}
	return res
}


func (this *NumArray) SumRange(i int, j int) int {
	if i == 0{
		return this.Sum[j]
	}else{
		return this.Sum[j] - this.Sum[i-1]
	}
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end


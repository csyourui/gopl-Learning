/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	M := make(map[int][]int)
	//seen := make(map[int]bool)
	innumber := make([]int, numCourses)
	var queue, res []int
	for _, pres := range prerequisites {
		//[pres[0], [pres[1]]
		//pres[1]------>pres[0]
		M[pres[1]] = append(M[pres[1]], pres[0])
		innumber[pres[0]]++
	}
	for i, v := range innumber {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0{
		front := queue[0]
		res = append(res, front)
		queue = queue[1:len(queue)]
		for _, val := range M[front] {
			innumber[val]--
			if innumber[val] == 0{
				queue = append(queue, val)
			}
		}
	}
	if len(res) == numCourses {return res}
	return nil
}
//2\n[[1,0],[0,1]]
// @lc code=end


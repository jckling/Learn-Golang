/*
 * @lc app=leetcode.cn id=697 lang=golang
 *
 * [697] Degree of an Array
 */

// @lc code=start
func findShortestSubArray(nums []int) int {
	pos := map[int][]int{}
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}
	degree := 0
	length := 0
	for i := range pos {
		if len(pos[i]) == degree {
			tmp := pos[i][len(pos[i])-1] - pos[i][0] + 1
			if length == 0 || length > tmp {
				length = tmp
			}
		} else if len(pos[i]) > degree {
			degree = len(pos[i])
			length = pos[i][len(pos[i])-1] - pos[i][0] + 1
		}
	}
	return length
}

// @lc code=end


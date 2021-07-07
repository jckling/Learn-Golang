/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */

// @lc code=start
func generate(numRows int) [][]int {
	var tmp int
	var arr [][]int
	for i := 0; i < numRows; i++ {
		row := []int{}
		for j := 0; j <= i; j++ {
			if j == 0 || i == 0 {
				tmp = 1
			} else {
				tmp = tmp * (i - j + 1) / j
			}
			row = append(row, tmp)
		}
		arr = append(arr, row)
	}
	return arr
}

// @lc code=end


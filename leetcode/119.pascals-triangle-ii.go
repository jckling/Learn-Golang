/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	var tmp int
	var arr []int

	for j := 0; j <= rowIndex; j++ {
		if j == 0 {
			tmp = 1
		} else {
			tmp = tmp * (rowIndex - j + 1) / j
		}
		arr = append(arr, tmp)
	}

	return arr
}

// @lc code=end


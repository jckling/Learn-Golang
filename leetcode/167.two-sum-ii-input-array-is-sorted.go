/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] Two Sum II - Input array is sorted
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	for i, x := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			if x+numbers[j] == target {
				return []int{i+1, j+1}
			}
		}
	}
	return nil
}
// @lc code=end


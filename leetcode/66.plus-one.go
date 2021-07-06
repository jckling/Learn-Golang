/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 < 10 {
			digits[i] = digits[i] + 1
			return digits
		}
		digits[i] = 0
	}

	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

// @lc code=end


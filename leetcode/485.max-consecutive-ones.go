/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] Max Consecutive Ones
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			cnt := 1
			for j := i + 1; j < len(nums); j++ {
				if nums[i] == nums[j] {
					cnt++
				} else {
					i = j
					break
				}
			}
			if cnt > max {
				max = cnt
			}
		}
	}
	return max
}

// @lc code=end

